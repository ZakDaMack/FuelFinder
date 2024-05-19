package main

import (
	"fmt"
	"log"
	"log/slog"
	"main/internal/api"
	"main/internal/env"
	"main/pkg/fuelfinder"
	"net/http"
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
	grpcHost := env.Get("GRPC_HOST", "localhost:50051")
	port := env.GetInt("PORT", 8080)

	client, err := fuelfinder.NewClient(grpcHost)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer client.Connection.Close()

	// set up http client, register routes
	gateway := api.NewGateway(&client.Commands)

	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	err = http.ListenAndServe(addr, gateway.Handler)
	if err != nil {
		log.Fatalf("error starting http server: %v", err)
	}
}
