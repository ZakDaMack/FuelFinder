package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/api/fueldata"
	"main/internal/env"
	"main/internal/scraper"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// get env vars
	grpcHost := env.Get("GRPC_HOST", "localhost:50051")

	// set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// close conn once done, gen client
	defer conn.Close()
	client := fueldata.NewFuelDataClient(conn)

	// read through links and then upload to system
	links, _ := scraper.GetTableLinks("https://www.gov.uk/guidance/access-fuel-price-data")
	for _, link := range links {
		data, err := scraper.ReadJsonFrom(link)
		fmt.Println("collected", len(data), "stations from", link)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		req := &fueldata.StationItems{Items: data}
		res, err := client.Upload(context.TODO(), req)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		fmt.Println("uploaded", res.Count, "stations")
	}
}
