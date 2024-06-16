package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kopicee/hdm-go/lib/services"
)

type API struct {
	router *gin.Engine
}

func NewAPI(hotelService services.HotelsService) API {
	hotels := NewHotelsController(hotelService)

	router := gin.Default()
	router.GET("/api/hotels", hotels.Find)

	return API{router}
}

func (a API) Listen(port int) error {
	binding := ":" + fmt.Sprint(port)

	return a.router.Run(binding)
}
