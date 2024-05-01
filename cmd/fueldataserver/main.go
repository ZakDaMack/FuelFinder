package main

import (
	"fmt"
	"log"
	"log/slog"
	"main/api/fueldata"
	"main/internal/env"
	"main/pkg/controllers"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	// set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// get env vars
	port := env.GetInt("PORT", 50051)
	mongoUri := env.Get("MONGO_URI", "mongodb://localhost:27017")

	slog.Info("setting up server...")
	slog.Info("gRPC port opened", "port", port)

	lis, err := net.Listen("tcp", fmt.Sprint(":", strconv.Itoa(port)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	slog.Info("connecting to mongo", "uri", mongoUri)

	// spool up grpc server
	grpcServer := grpc.NewServer()
	fds, err := controllers.NewFuelDataServer("ofd", mongoUri)

	// on server start, ensrue that the indexes are working, or the queries wont work
	fds.EnsureIndexes()

	if err != nil {
		log.Fatalf("failed to setup fueldata server: %v", err)
	}
	fueldata.RegisterFuelDataServer(grpcServer, fds)

	slog.Info("serving server", "uri", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
