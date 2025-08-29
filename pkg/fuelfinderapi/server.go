package fuelfinderapi

import "github.com/gin-gonic/gin"

func NewServer(params *FuelFinderAPIParams) *gin.Engine {
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")

	return router
}
