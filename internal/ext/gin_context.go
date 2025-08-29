package ext

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryFloat(ctx *gin.Context, key string) float64 {
	val, ok := ctx.GetQuery(key)
	if !ok {
		return 0
	}

	fl, _ := strconv.ParseFloat(val, 64)
	return fl
}
