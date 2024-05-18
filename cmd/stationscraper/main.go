package main

import (
	"context"
	"log/slog"
	"main/api/fueldata"
	"main/internal/env"
	"main/internal/sanitiser"
	"main/internal/scraper"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Job struct {
	url              string
	stations         []*fueldata.StationItem
	uploadedStations int32
}

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
	url := "https://www.gov.uk/guidance/access-fuel-price-data"
	links, err := scraper.GetTableLinks(url)
	if err != nil {
		slog.Error("could not scrape fuel price data", "url", url, "error", err)
		return
	}

	slog.Info("fetched companies from gov website", "count", len(links))

	// GO CHANNEL PIPELINE
	// stage 1: convert array to jobs
	jobs := createJobsChannel(links)
	// stage 2: collect stations from jobs
	res := fetchData(jobs)
	// stage 3: filter data
	res2 := filterData(res)
	// stage 4: sanitize data
	res3 := sanitizeData(res2)
	// stage 5: save
	res4 := uploadData(res3, client)
	// finally print done line
	for j := range res4 {
		slog.Info("finished job for station url", "count", j.uploadedStations, "url", j.url)
	}
}

func createJobsChannel(items []string) <-chan Job {
	out := make(chan Job)
	go func(arr []string) {
		for _, i := range arr {
			out <- Job{
				url: i,
			}
		}
		close(out)
	}(items)

	return out
}

func fetchData(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range in {
			data, err := scraper.ReadJsonFrom(job.url)
			if err != nil {
				slog.Error("could not read json", "url", job.url, "error", err)
				continue
			}

			job.stations = data
			out <- job
		}
		close(out)
	}()

	return out
}

func filterData(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range in {
			// Sanitise the data
			sanitisedStations := make([]*fueldata.StationItem, 0)
			for _, station := range job.stations {
				if sanitiser.IsValidItem(station) {
					sanitisedStations = append(sanitisedStations, station)
				}
			}
			job.stations = sanitisedStations
			out <- job
		}
		close(out)
	}()

	return out
}

func sanitizeData(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range in {
			// Sanitise the data
			for _, station := range job.stations {
				sanitiser.CleanStationItem(station)
			}
			out <- job
		}
		close(out)
	}()

	return out
}

func uploadData(in <-chan Job, client fueldata.FuelDataClient) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range in {
			req := &fueldata.StationItems{Items: job.stations}
			res, err := client.Upload(context.TODO(), req)
			if err != nil {
				slog.Error("could not upload station items", "url", job.url, "error", err)
				continue
			}
			job.uploadedStations = res.Count
		}
		close(out)
	}()
	return out
}
