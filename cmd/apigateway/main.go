package main

import (
	"fmt"
	"log"
	"log/slog"
	"main/api/fueldata"
	"main/internal/env"
	"main/pkg/controllers"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// get env vars
	grpcHost := env.Get("GRPC_HOST", "localhost:50051")
	port := env.GetInt("PORT", 8080)

	// set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// set up grpc client
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// close conn once done, gen client
	defer conn.Close()
	client := fueldata.NewFuelDataClient(conn)

	// set up http client, register routes
	r := gin.Default()
	gw := &controllers.Gateway{Client: &client}
	r.GET("/", gw.GetStations)
	r.GET("/ping", gw.GetPing)

	r.Run(fmt.Sprintf(":%s", strconv.Itoa(port)))
}
