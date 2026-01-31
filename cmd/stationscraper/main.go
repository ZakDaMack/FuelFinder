package main

import (
	"context"
	"log/slog"
	"main/internal/database"
	"main/internal/env"
	"os/signal"
	"syscall"
	"time"

	"main/internal/scraper"
	"os"
)

func main() {
	// get env vars
	url := "https://www.gov.uk/guidance/access-fuel-price-data" // fixed for now
	mongoUri := env.Get("MONGO_URI", "mongodb://localhost:27017")
	dbName := "ofd"
	interval := env.GetInt("INTERVAL", 1)
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

	// connect to mongo
	conn, err := database.NewMongoConnection(mongoUri, dbName)
	if err != nil {
		slog.Error("error connecting to mongo", "error", err)
		return
	}

	defer conn.Close()

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

			// check if files already exists, if so, return
			exists, err := conn.Exists(job.Stations[0].CreatedAt, job.Stations[0].SiteId)
			if err != nil {
				return nil, err
			}

			if exists {
				res.Count = 0
				return res, nil
			}

			writeRes, err := conn.Write(items.Items)

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
