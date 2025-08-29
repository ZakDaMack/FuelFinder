package controller

import (
	"main/internal/model"
	"main/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StationController interface {
	Index(ctx *gin.Context)
}

type stationController struct {
	stationsRepo repository.StationRepository
}

func NewStationController(db *mongo.Database) StationController {
	return &stationController{
		stationsRepo: repository.NewMongoStationRepo(db),
	}
}

func (c *stationController) Index(ctx *gin.Context) {
	lat := ctx.GetFloat64("latitude")
	long := ctx.GetFloat64("longitude")
	radius := ctx.GetInt("radius")
	brandQuery := ctx.GetStringSlice("brands")
	fueltypes := ctx.GetStringSlice("fueltypes")

	if radius < 1 || radius > 20 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Radius must be between 1 and 20",
		})
		return
	}

	stations, err := c.stationsRepo.GetStations(ctx, lat, long, radius, brandQuery, fueltypes)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, stations)
}
