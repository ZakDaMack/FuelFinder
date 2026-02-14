package controller

import (
	"errors"
	"fmt"
	"main/internal/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getDelimitedQueryParam(ctx *gin.Context, paramName string) []string {
	paramValue := ctx.Query(paramName)
	if paramValue == "" {
		return []string{}
	}

	return strings.Split(paramValue, ",")
}

func getDelimitedFloatQueryParam(ctx *gin.Context, paramName string) ([]float64, error) {
	paramValue := ctx.Query(paramName)
	if paramValue == "" {
		return []float64{}, nil
	}

	strValues := strings.Split(paramValue, ",")
	floatValues := make([]float64, 0, len(strValues))
	for _, strVal := range strValues {
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float value for query param %s: %v", paramName, err)
		}
		floatValues = append(floatValues, floatVal)
	}
	return floatValues, nil
}

// handles app defined errors and aborts if not nil
func handleError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	// if our custom error type, extract status and message
	var appErr *model.ErrorResponse
	if errors.As(err, &appErr) {
		ctx.AbortWithStatusJSON(appErr.Status, appErr)
		return true
	}

	// else just return 500 with generic message
	ctx.AbortWithError(500, &model.ErrorResponse{
		Code:    "internal_error",
		Message: err.Error(),
	})
	return true
}
