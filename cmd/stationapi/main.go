package main

import (
	"log/slog"
	"main/internal/env"
	"main/pkg/stationapi"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	// get env vars
	port := env.GetInt("PORT", 8080)
	postgisDSN := env.Get("POSTGIS_DSN", "postgres://localhost:5432/fuelfinder?sslmode=disable")
	debugMode := env.ExistsAndNotFalse("DEBUG_MODE")

	// set up logging
	options := &slog.HandlerOptions{}
	if debugMode {
		options.Level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, options))
	slog.SetDefault(logger)

	// listen for signal kill
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	slog.Info("setting up server...")
	server := stationapi.NewServer(&stationapi.StationAPIParams{
		Port:       port,
		PostgisDSN: postgisDSN,
	})

	server.Start()
	slog.Info("server started", "port", strconv.Itoa(port))

	// wait for SIGKILL
	s := <-sigChan
	slog.Info("app killed through signal", "signal", s.String())
	server.Stop()
}
