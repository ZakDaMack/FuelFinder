package main

import (
	"context"
	"log"
	"log/slog"
	"main/internal/env"
	"time"

	fuelfinderproto "main/api/fuelfinder"
	"main/internal/scraper"
	"main/pkg/fuelfinder"
	"os"
)

func main() {
	// set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	slog.Info("scraper started")

	// get env vars
	slog.Debug("getting env vars")
	url := "https://www.gov.uk/guidance/access-fuel-price-data" // fixed for now
	grpcHost := env.Get("GRPC_HOST", "localhost:50051")
	interval := env.GetInt("INTERVAL", 1)
	slog.Debug("got env vars", "host", grpcHost, "interval", interval)

	// create client
	client, err := fuelfinder.NewClient(grpcHost)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer client.Connection.Close()

	// make an outbound channel for new stationdata
	data := make(chan scraper.Job)

	slog.Debug("creating ticker")
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)

	for {
		select {
		case t := <-ticker.C:
			slog.Info("job triggered", "ticker_time", t)
			go scraper.NewScraper(url, data)
		case job := <-data:
			slog.Info("job passed to upload", "url", job.Url, "items", len(job.Stations))
			uploaded, err := client.Commands.Upload(context.TODO(), &fuelfinderproto.StationItems{
				Items: job.Stations,
			})
			if err != nil {
				slog.Error("error uploading station data", "url", job.Url, "error", err)
			} else {
				slog.Info("finished job for station url", "url", job.Url, "count", uploaded.Count)
			}
		}
	}
}
