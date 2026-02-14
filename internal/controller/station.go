package controller

import (
	"main/internal/service"

	"github.com/gin-gonic/gin"
)

type StationController interface {
	GetQuery(ctx *gin.Context)
	GetBrands(ctx *gin.Context)
	GetStation(ctx *gin.Context)
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
	if handleError(ctx, err) {
		return
	}

	// get brands and parse into slice
	brandQuery := getDelimitedQueryParam(ctx, "brands")

	// get fuel types and parse into slice
	fuelQuery := getDelimitedQueryParam(ctx, "fueltypes")

	stations, err := c.stationService.GetStations(ctx, coords, brandQuery, fuelQuery)
	if handleError(ctx, err) {
		return
	}

	ctx.JSON(200, stations)
}

func (c *stationController) GetBrands(ctx *gin.Context) {
	brands, err := c.stationService.GetBrands(ctx)
	if handleError(ctx, err) {
		return
	}

	ctx.JSON(200, brands)
}

func (c *stationController) GetStation(ctx *gin.Context) {
	siteID := ctx.Param("id")
	stations, err := c.stationService.GetStationsBySiteID(ctx, siteID)
	if handleError(ctx, err) {
		return
	}

	ctx.JSON(200, stations)
}
