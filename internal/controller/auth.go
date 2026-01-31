package controller

import (
	"main/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx gin.Context)
	Logout(ctx gin.Context)
	Register(ctx gin.Context)
	RefreshToken(ctx gin.Context)
	WellKnown(ctx gin.Context)
}

type authController struct {
	authService AuthService
}

func NewAuthController(authService AuthService) AuthController {
	return &authController{
		authService: service.NewAuthService(service.JWTOptions{
			Issuer: "FuelFinder",
		}),
	}
}

func (c *authController) WellKnown(ctx gin.Context) {
	ctx.JSON(200, gin.H{
		"issuer": "FuelFinder",
	})
}
