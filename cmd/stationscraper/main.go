package main

import (
	"context"
	"log"
	"log/slog"
	"main/internal/env"
	"os/signal"
	"syscall"
	"time"

	fuelfinderproto "main/api/fuelfinder"
	"main/internal/scraper"
	"main/pkg/fuelfinder"
	"os"
)

func main() {
	// get env vars
	url := "https://www.gov.uk/guidance/access-fuel-price-data" // fixed for now
	grpcHost := env.Get("GRPC_HOST", "localhost:50051")
	interval := env.GetInt("INTERVAL", 1)
	debugMode := env.ExistsAndNotFalse("DEBUG_MODE")

	// set up logging
	options := &slog.HandlerOptions{}
	if debugMode {
		options.Level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, options))
	slog.SetDefault(logger)

	slog.Debug("got env vars", "host", grpcHost, "interval", interval)
	slog.Info("scraper started")

	// create client
	client, err := fuelfinder.NewClient(grpcHost)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer client.Connection.Close()

	ctx := context.Background()
	ctx.
		ch := <-ctx.Done()
	// ctx

	// listen for signal kill
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// make an outbound channel for new stationdata
	data := make(chan scraper.Job)

	slog.Debug("creating ticker")
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)

outer:
	for {
		select {
		case t := <-ticker.C:
			slog.Info("job triggered", "ticker_time", t)
			go scraper.NewScraper(url, data)
		case job := <-data:
			slog.Debug("job passed to upload", "url", job.Url, "items", len(job.Stations))
			uploaded, err := client.Commands.Upload(context.TODO(), &fuelfinderproto.StationItems{
				Items: job.Stations,
			})
			if err != nil {
				slog.Error("error uploading station data", "url", job.Url, "error", err)
			} else {
				slog.Info("finished job for station url", "url", job.Url, "uploaded", uploaded.Count)
			}
		case s := <-sigChan:
			slog.Info("app killed through signal", "signal", s.String())
			break outer
		}
	}
}
