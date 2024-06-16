package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type baseController struct{}

func (c baseController) respondJSON(ctx *gin.Context, statusCode int, result any) {
	ctx.JSON(statusCode, result)
}

func (c baseController) Resolve(ctx *gin.Context, result any) {
	c.respondJSON(ctx, 200, result)
}

func (c baseController) Reject(ctx *gin.Context, err error) {
	body := gin.H{
		"error": err.Error(),
	}

	var statusCode int
	switch {
	case errors.Is(err, errBadQuery):
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	c.respondJSON(ctx, statusCode, body)
}

var errBadQuery = errors.New("error parsing query")

func getQueryAs[T any](ctx *gin.Context, key string, parser func(string) (T, error)) ([]T, error) {
	values := make([]T, 0)

	for _, stringVal := range ctx.QueryArray(key) {
		var value T
		var err error

		if parser == nil {
			value, _ = any(stringVal).(T)
		} else {
			value, err = parser(stringVal)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", errBadQuery, err)
			}
		}
		values = append(values, value)
	}
	return values, nil
}
