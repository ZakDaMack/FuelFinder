package controllers

import (
	"context"
	"main/api/fueldata"
	"strconv"
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
	lat, _ := strconv.ParseFloat(c.Query("latitude"), 32)
	long, _ := strconv.ParseFloat(c.Query("longitude"), 32)
	radius, _ := strconv.ParseFloat(c.Query("radius"), 32)

	if radius < 1 || radius > 20 {
		c.JSON(400, gin.H{
			"message": "Radius must be between 1 and 20",
		})
		return
	}

	service := *g.Client
	val, err := service.QueryArea(context.TODO(), &fueldata.Geofence{
		Latitude:  float32(lat),
		Longitude: float32(long),
		Radius:    float32(radius),
	})

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, val.Items)
}

// func (g *Gateway) GetBrands(c *gin.Context) {
// 		service := *g.Client
// 	val, err := service.QueryArea()(context.TODO(), &fueldata.Geofence{
// 		Latitude:  float32(lat),
// 		Longitude: float32(long),
// 		Radius:    float32(radius),
// 	})

// 	if err != nil {
// 		c.JSON(500, err)
// 		return
// 	}

// 	c.JSON(200, val.Items)
// }
