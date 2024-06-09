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
	debugMode := env.GetBool("DEBUG_MODE", false)
	immediate := true

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
			items := &fuelfinderproto.StationItems{Items: job.Stations}
			uploaded, err := client.Commands.Upload(context.TODO(), items)
			if err != nil {
				slog.Error("error uploading station data", "url", job.Url, "error", err)
			} else {
				slog.Info("finished job for station url", "url", job.Url, "uploaded", uploaded.Count)
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
