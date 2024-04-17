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
	// get env vars
	port := env.GetInt("PORT", 50051)
	mongoUri := env.Get("MONGO_URI", "mongodb://localhost:27017")

	slog.Info("setting up server...")

	// set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("gRPC port opened", "port", port)

	lis, err := net.Listen("tcp", fmt.Sprint(":", strconv.Itoa(port)))
	if err != nil {
		log.Fatalf("FATAL: failed to listen: %v", err)
	}

	slog.Info("connecting to mongo", "uri", mongoUri)

	// spool up grpc server
	grpcServer := grpc.NewServer()
	fds, err := controllers.NewFuelDataServer("ofd", mongoUri)
	if err != nil {
		log.Fatalf("FATAL: failed to setup fueldata server: %v", err)
	}
	fueldata.RegisterFuelDataServer(grpcServer, fds)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("FATAL: failed to serve: %v", err)
	}
}
