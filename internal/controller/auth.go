package controller

import (
	"main/internal/model"
	"main/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	PostLogin(ctx *gin.Context)
	PostLogout(ctx *gin.Context)
	PostRegister(ctx *gin.Context)
	PostRefreshToken(ctx *gin.Context)
	GetWellKnown(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(supplier service.Supplier) AuthController {
	return &authController{
		authService: supplier.GetAuthService(),
	}
}

func (c *authController) PostLogin(ctx *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindBodyWithJSON(&loginRequest); err != nil {
		return
	}

	jwtToken, refreshToken, err := c.authService.Login(ctx, loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(400, model.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)
	ctx.JSON(200, gin.H{
		"token":         jwtToken,
		"refresh_token": refreshToken,
	})
}

func (c *authController) PostLogout(ctx *gin.Context) {
	ctx.Status(201)
}

func (c *authController) PostRegister(ctx *gin.Context) {
	var registerRequest struct {
		Email           string `json:"email" binding:"required,email"`
		Password        string `json:"password" binding:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
		FirstName       string `json:"first_name" binding:"required"`
		LastName        string `json:"last_name" binding:"required"`
	}
	if err := ctx.ShouldBindBodyWithJSON(&registerRequest); err != nil {
		return
	}

	jwtToken, refreshToken, err := c.authService.Register(
		ctx,
		registerRequest.Email,
		registerRequest.Password,
		registerRequest.ConfirmPassword,
		registerRequest.FirstName,
		registerRequest.LastName,
	)
	if err != nil {
		ctx.JSON(400, model.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)
	ctx.JSON(200, gin.H{
		"token":         jwtToken,
		"refresh_token": refreshToken,
	})
}

func (c *authController) PostRefreshToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(400, model.ErrorResponse{Message: "refresh token missing"})
		return
	}

	jwtToken, err := c.authService.RefreshToken(ctx, cookie)
	if err != nil {
		ctx.JSON(400, model.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"token": jwtToken,
	})
}

func (c *authController) GetWellKnown(ctx *gin.Context) {
	jwks, err := c.authService.GetWellKnown(ctx)
	if err != nil {
		ctx.JSON(500, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(200, jwks)
}
