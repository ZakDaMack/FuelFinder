package main

import (
	"context"
	"log/slog"
	"main/api/fueldata"
	"main/internal/env"
	"main/internal/scraper"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("scraper started")

	// get env vars
	slog.Debug("getting env vars")
	grpcHost := env.Get("GRPC_HOST", "localhost:50051")
	interval := env.GetInt("INTERVAL", 1)
	slog.Debug("got env vars", "host", grpcHost, "interval", interval)

	slog.Debug("creating ticker")
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	for t := range ticker.C {
		slog.Info("job triggered", "ticker_time", t)
		scrapeData(grpcHost)
	}
	slog.Info("app exiting")
}

func scrapeData(grpcHost string) {

	// dial the grpc server host
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("could not connect to grpc", "host", grpcHost, "error", err)
		return
	}

	// close conn once done, gen client
	defer conn.Close()
	client := fueldata.NewFuelDataClient(conn)

	// read through links and then upload to system
	links, err := scraper.GetTableLinks("https://www.gov.uk/guidance/access-fuel-price-data")
	if err != nil {
		slog.Error("could not scrape fuel price data", "url", "https://www.gov.uk/guidance/access-fuel-price-data", "error", err)
		return
	}
	for _, link := range links {
		slog.Info("reading json", "url", link)

		data, err := scraper.ReadJsonFrom(link)
		if err != nil {
			slog.Error("could not read json", "url", link, "error", err)
			return
		}

		slog.Debug("collected station data", "link", link, "stations", len(data))

		req := &fueldata.StationItems{Items: data}
		res, err := client.Upload(context.TODO(), req)
		if err != nil {
			slog.Error("could not read json", "url", link, "error", err)
			return
		}

		slog.Debug("uploaded station data", "stations", res.Count)
	}
}
