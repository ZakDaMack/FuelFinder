package controller

import (
	"main/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	GetPing(ctx *gin.Context)
}

type healthController struct {
	startupTime time.Time
}

func NewHealthController() HealthController {
	return &healthController{
		startupTime: time.Now(),
	}
}

func (c *healthController) GetPing(ctx *gin.Context) {
	ctx.JSON(200, &model.PingResponse{
		Message:     "OK",
		Uptime:      int64(time.Since(c.startupTime).Seconds()),
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
	})
}
