package controller

import (
	"log/slog"
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StationController interface {
	GetQuery(ctx *gin.Context)
	GetBrands(ctx *gin.Context)
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
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// get brands and parse into slice
	brandQuery := getDelimitedQueryParam(ctx, "brands")

	// get fuel types and parse into slice
	fuelQuery := getDelimitedQueryParam(ctx, "fueltypes")

	slog.Info("got query", "coords", coords, "brands", brandQuery, "fueltypes", fuelQuery)
	stations, err := c.stationService.GetStations(ctx, coords, brandQuery, fuelQuery)
	if err != nil {
		if err.Error() == service.IncorrectCoordsError || err.Error() == service.DistanceTooGreatError {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	slog.Info("got stations", "len", len(stations))
	ctx.JSON(200, stations)
}

func (c *stationController) GetBrands(ctx *gin.Context) {
	brands, err := c.stationService.GetBrands(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, brands)
}
