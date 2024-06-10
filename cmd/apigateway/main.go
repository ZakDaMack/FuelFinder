package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/internal/api"
	"main/internal/env"
	"main/pkg/fuelfinder"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// get env vars
	grpcHost := env.Get("GRPC_HOST", "localhost:50051")
	port := env.GetInt("PORT", 8080)
	debugMode := env.ExistsAndNotFalse("DEBUG_MODE")

	// set up logging
	options := &slog.HandlerOptions{}
	if debugMode {
		options.Level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, options))
	slog.SetDefault(logger)

	client, err := fuelfinder.NewClient(grpcHost)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer client.Connection.Close()

	// set up http client, register routes
	addr := fmt.Sprintf(":%d", port)
	server := api.NewGateway(&client.Commands, addr)

	// listen for signal kill
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting http server: %v", err)
		}
	}()

	// wait for SIGKILL
	s := <-sigChan
	slog.Info("app killed through signal", "signal", s.String())

	//shutdown server with 5 second limit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("error shutting down http server: %v", err)
	}

}
