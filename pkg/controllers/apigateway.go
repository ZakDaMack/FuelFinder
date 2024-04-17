package controllers

import (
	"context"
	"main/api/fueldata"
	"time"

	"github.com/gin-gonic/gin"
)

type Gateway struct {
	Client *fueldata.FuelDataClient
}

func (*Gateway) GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":   "pong",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (g *Gateway) GetStations(c *gin.Context) {
	// c.GetQueryMap()
	service := *g.Client
	val, _ := service.QueryArea(context.TODO(), &fueldata.Geofence{
		Latitude:  51.795971,
		Longitude: -0.078880,
		Radius:    3,
	})
	c.JSON(200, val.Items)
}
