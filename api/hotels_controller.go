package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kopicee/hdm-go/lib/services"
)

type hotelsController struct {
	baseController
	hotels services.HotelsService
}

func NewHotelsController(svc services.HotelsService) hotelsController {
	return hotelsController{baseController{}, svc}
}

func (ctrl hotelsController) Find(ctx *gin.Context) {
	hotelIDs, err := getQueryAs[string](ctx, "id", nil)
	if err != nil {
		ctrl.Reject(ctx, err)
		return
	}

	destinationIDs, err := getQueryAs[int](ctx, "destination", strconv.Atoi)
	if err != nil {
		ctrl.Reject(ctx, err)
		return
	}

	found, err := ctrl.hotels.Find(hotelIDs, destinationIDs)
	if err != nil {
		ctrl.Reject(ctx, err)
		return
	}

	ctrl.Resolve(ctx, found)
}
