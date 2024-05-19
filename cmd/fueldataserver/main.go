package main

import (
	"fmt"
	"log"
	"log/slog"
	"main/internal/env"
	"main/pkg/fuelfinder"
	"net"
	"os"
	"strconv"
)

func main() {
	// set up logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// get env vars
	port := env.GetInt("PORT", 50051)
	mongoUri := env.Get("MONGO_URI", "mongodb://localhost:27017")

	slog.Info("setting up server...")
	slog.Info("gRPC port opened", "port", port)

	addr := fmt.Sprint(":", strconv.Itoa(port))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	slog.Info("connecting to mongo", "uri", mongoUri)

	// spool up grpc server
	grpcServer, err := fuelfinder.NewGrpcServer("ofd", mongoUri)
	if err != nil {
		log.Fatalf("failed to setup fueldata server: %v", err)
	}

	slog.Info("starting server", "uri", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
