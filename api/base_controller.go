package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrBadQuery = errors.New("error parsing query")

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
	case errors.Is(err, ErrBadQuery):
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	c.respondJSON(ctx, statusCode, body)
}

func newBadQuery(keyName string) error {
	return fmt.Errorf("%w '%s'", ErrBadQuery, keyName)
}

func getQueryAs[T any](ctx *gin.Context, key string, parser func(string) (T, error)) ([]T, error) {
	values := make([]T, 0)

	for _, stringVal := range ctx.QueryArray(key) {
		var value T
		var err error
		var ok bool

		if stringVal == "" {
			continue
		}

		if parser == nil {
			value, ok = any(stringVal).(T)
			if !ok {
				return nil, fmt.Errorf("%w: not a %T", newBadQuery(key), value)
			}
		} else {
			value, err = parser(stringVal)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", newBadQuery(key), err)
			}
		}
		values = append(values, value)
	}
	return values, nil
}
