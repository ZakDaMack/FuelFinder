package stationapi

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/internal/controller"
	"main/internal/database"
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(params *StationAPIParams) *Server {
	router := gin.Default()
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", params.Port),
		Handler: router,
	}

	// set middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// set up db, services and controllers
	database, err := database.MakePostgresDB(params.PostgisDSN)
	if err != nil {
		log.Fatal("failed to connect to database", "err", err)
	}

	stationSvc := service.NewStationService(database)
	stationController := controller.NewStationController(stationSvc)
	healthController := controller.NewHealthController()

	// Health routes
	router.GET("/health/ping", healthController.GetPing)

	// Station routes
	router.GET("/stations/query", stationController.GetQuery)
	router.GET("/stations/brands", stationController.GetBrands)
	router.GET("/stations/:id", stationController.GetStation)

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() {
	// start server in goroutine
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			slog.Error("encountered error with http server", "err", err)
		}
	}()
}

func (s *Server) Stop() {
	s.httpServer.Shutdown(context.Background())
}
