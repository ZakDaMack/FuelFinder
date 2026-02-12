package controller

import (
	"log/slog"
	"main/internal/service"

	"github.com/gin-gonic/gin"
)

type StationController interface {
	GetQuery(ctx *gin.Context)
	GetBrands(ctx *gin.Context)
	GetStations(ctx *gin.Context)
}

type stationController struct {
	stationService service.StationService
}

func NewStationControllerFromSupplier(supplier service.Supplier) StationController {
	return NewStationController(supplier.GetStationService())
}

func NewStationController(stationSvc service.StationService) StationController {
	return &stationController{
		stationService: stationSvc,
	}
}

func (c *stationController) GetQuery(ctx *gin.Context) {
	// get coords from query params and parse into slice of floats
	coords, err := getDelimitedFloatQueryParam(ctx, "coords")
	handleError(ctx, err)

	// get brands and parse into slice
	brandQuery := getDelimitedQueryParam(ctx, "brands")

	// get fuel types and parse into slice
	fuelQuery := getDelimitedQueryParam(ctx, "fueltypes")

	slog.Info("got query", "coords", coords, "brands", brandQuery, "fueltypes", fuelQuery)
	stations, err := c.stationService.GetStations(ctx, coords, brandQuery, fuelQuery)
	handleError(ctx, err)

	slog.Info("got stations", "len", len(stations))
	ctx.JSON(200, stations)
}

func (c *stationController) GetBrands(ctx *gin.Context) {
	brands, err := c.stationService.GetBrands(ctx)
	handleError(ctx, err)

	ctx.JSON(200, brands)
}

func (c *stationController) GetStations(ctx *gin.Context) {
	siteID := ctx.Param("site_id")
	stations, err := c.stationService.GetStationsBySiteID(ctx, siteID)
	handleError(ctx, err)

	ctx.JSON(200, stations)
}
