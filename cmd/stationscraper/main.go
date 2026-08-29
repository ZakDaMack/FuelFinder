package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"main/internal/dao"
	"main/internal/database"
	"main/internal/env"
	"main/internal/govapi"
	"main/internal/repository"
	"main/internal/service"
)

func main() {
	// get env vars
	databaseDSN := env.Get("POSTGIS_DSN", "postgres://localhost:5432/ofd?sslmode=disable")
	apiURL := env.Get("GOV_API_URL", "https://www.fuel-finder.service.gov.uk")
	clientID := env.Get("GOV_API_CLIENT_ID", "")
	clientSecret := env.Get("GOV_API_CLIENT_SECRET", "")

	interval := env.GetInt("SCRAPER_INTERVAL", 15)
	debugMode := env.ExistsAndNotFalse("DEBUG_MODE")
	immediate := env.ExistsAndNotFalse("IMMEDIATE")

	// set up logging
	options := &slog.HandlerOptions{}
	if debugMode {
		options.Level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, options))
	slog.SetDefault(logger)

	slog.Info("scraper started")
	slog.Info("getting envs", "url", apiURL, "dsn", databaseDSN, "debug", debugMode, "interval", interval, "immediate", immediate)

	if clientID == "" || clientSecret == "" {
		slog.Error("GOV_API_CLIENT_ID and GOV_API_CLIENT_SECRET must be set")
		return
	}

	// connect to database
	db, err := database.MakePostgresDB(databaseDSN)
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		return
	}

	// ensure db closed on termination of session
	defer database.ClosePostgresDB(db)

	// set up repository connections
	stationRepo := repository.NewStationRepo(db)
	priceRepo := repository.NewStationPriceRepo(db)

	// set up gov api service
	govAPI := service.NewGovAPIService(clientID, clientSecret, apiURL)

	// listen for signal kill
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// create ticker depending on if immediate or not
	slog.Debug("creating ticker")
	var dur time.Duration
	if immediate {
		dur = 1 * time.Second
	} else {
		dur = time.Duration(interval) * time.Minute
	}
	ticker := time.NewTicker(dur)

	// create cancel context
	ctx, cancel := context.WithCancel(context.Background())

	// fetch everything on the first job, then only changes since the last success
	var lastRun *time.Time

outer:
	for {
		select {
		case t := <-ticker.C:
			if immediate {
				ticker.Stop()
			}
			slog.Info("job triggered", "ticker_time", t)

			lastRun = runJob(ctx, govAPI, stationRepo, priceRepo, lastRun)

			// cancel context after first job completes, if immediate is true this will kill the app
			if immediate {
				cancel()
			}

		case <-ctx.Done():
			slog.Info("app killed after completion")
			break outer

		case s := <-sigChan:
			slog.Info("app killed through signal", "signal", s.String())
			break outer
		}
	}
}

func runJob(ctx context.Context, govAPI service.GovAPIService, stationRepo repository.StationRepository, priceRepo repository.StationPriceRepository, lastRun *time.Time) *time.Time {
	started := time.Now().UTC()

	stations, err := govAPI.GetAllStations(lastRun)
	if err != nil {
		slog.Error("error fetching stations", "error", err)
		return lastRun
	}

	// upsert stations into database, keyed by node id
	errs := make([]error, 0)
	for _, station := range stations {
		if err := stationRepo.Upsert(ctx, toStation(station)); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		slog.Error("error uploading station data", "successes", len(stations)-len(errs), "failures", len(errs), "errors", errs)
	} else {
		slog.Info("finished station job", "uploaded", len(stations))
	}

	// fetch fuel prices and insert them as new price records
	nodePrices, err := govAPI.GetAllFuelPrices(lastRun)
	if err != nil {
		slog.Error("error fetching fuel prices", "error", err)
		return lastRun
	}

	prices := make([]*dao.StationPrice, 0, len(nodePrices))
	for i := range nodePrices {
		if price := toStationPrice(&nodePrices[i]); price != nil {
			prices = append(prices, price)
		}
	}

	if err := priceRepo.Insert(ctx, prices); err != nil {
		slog.Error("error uploading fuel price data", "error", err)
		return lastRun
	}

	slog.Info("finished price job", "nodes", len(nodePrices), "price_records", len(prices))

	// only advance the watermark once a full sync has succeeded
	if len(errs) > 0 {
		return lastRun
	}
	return &started
}

func toStation(node govapi.StationNode) *dao.Station {
	phone := ""
	if node.PublicPhoneNumber != nil {
		phone = *node.PublicPhoneNumber
	}

	addressParts := []string{node.Location.AddressLine1}
	if node.Location.AddressLine2 != nil && *node.Location.AddressLine2 != "" {
		addressParts = append(addressParts, *node.Location.AddressLine2)
	}
	addressParts = append(addressParts, node.Location.City)

	return &dao.Station{
		NodeID:   node.NodeID,
		Name:     node.TradingName,
		Brand:    node.BrandName,
		Address:  strings.Join(addressParts, ", "),
		Postcode: node.Location.Postcode,
		Phone:    phone,
		Location: dao.GeometryPoint{
			X: node.Location.Longitude,
			Y: node.Location.Latitude,
		},
	}
}

func toStationPrice(node *govapi.NodeFuelPrice) *dao.StationPrice {
	price := &dao.StationPrice{NodeID: node.NodeID}

	var effectiveAt time.Time
	matched := false
	for _, fuel := range node.FuelPrices {
		val := sql.NullFloat64{Float64: fuel.Price, Valid: true}

		switch strings.ToUpper(fuel.FuelType) {
		case "E5":
			price.E5 = val
		case "E10":
			price.E10 = val
		case "B7_STANDARD":
			price.B7Standard = val
		case "B7_PREMIUM":
			price.B7Premium = val
		default:
			// B10 and HVO have no columns yet
			slog.Debug("skipping unsupported fuel type", "node_id", node.NodeID, "fuel_type", fuel.FuelType)
			continue
		}
		matched = true

		// the record is effective from its latest price change
		if ts, err := time.Parse(time.RFC3339, fuel.PriceChangeEffectiveTimestamp); err == nil && ts.After(effectiveAt) {
			effectiveAt = ts
		}
	}

	if !matched {
		return nil
	}

	if effectiveAt.IsZero() {
		effectiveAt = time.Now().UTC()
	}
	price.EffectiveAt = effectiveAt

	return price
}
