package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"main/internal/dao"
	"main/internal/database"
	"main/internal/env"
	"main/internal/repository"
	"main/internal/sanitiser"
	"main/internal/scraper"
)

func main() {
	// get env vars
	url := env.Get("SCRAPER_URL", "https://www.gov.uk/guidance/access-fuel-price-data")
	databaseDSN := env.Get("POSTGIS_DSN", "postgres://localhost:5432/ofd?sslmode=disable")

	interval := env.GetInt("SCRAPER_INTERVAL", 1)
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

	slog.Info("getting envs", "url", url, "dsn", databaseDSN, "debug", debugMode)

	// connect to database
	db, err := database.MakePostgresDB(databaseDSN)
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		return
	}

	defer database.ClosePostgresDB(db)

	// set up repository connection
	repo := repository.NewStationRepo(db)

	// listen for signal kill
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// make an outbound channel for new stationdata
	data := make(chan scraper.Job)
	defer close(data)

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

outer:
	for {
		select {
		case t := <-ticker.C:
			if immediate {
				ticker.Stop()
			}
			slog.Info("job triggered", "ticker_time", t)
			go func(cncl context.CancelFunc) {
				scraper.NewScraper(url, data)
				cncl()
			}(cancel)

		case job := <-data:
			slog.Debug("job passed to upload", "url", job.Url, "items", len(job.Stations))

			// insert stations into database
			errs := make([]error, 0)
			for _, station := range job.Stations {
				err = repo.Insert(ctx, &dao.Station{
					SiteID:   station.SiteId,
					Brand:    station.Brand,
					Address:  station.Address,
					Postcode: station.Postcode,
					Location: dao.GeometryPoint{
						X: sanitiser.ToFloat(station.Location.Longitude),
						Y: sanitiser.ToFloat(station.Location.Latitude),
					},
					E5:        sql.NullFloat64{Float64: float64(station.Prices.E5), Valid: true},
					E10:       sql.NullFloat64{Float64: float64(station.Prices.E10), Valid: true},
					B7:        sql.NullFloat64{Float64: float64(station.Prices.B7), Valid: true},
					SDV:       sql.NullFloat64{Float64: float64(station.Prices.SDV), Valid: true},
					CreatedAt: *station.CreatedAt,
				})

				if err != nil {
					errs = append(errs, err)
				}
			} // end for

			if len(errs) > 0 {
				slog.Error("error uploading station data", "url", job.Url, "successes", len(job.Stations)-len(errs), "failures", len(errs), "errors", errs)
			} else {
				slog.Info("finished job for station url", "url", job.Url, "uploaded", len(job.Stations))
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
