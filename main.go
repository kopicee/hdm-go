package main

import (
	"github.com/kopicee/hdm-go/api"
	"github.com/kopicee/hdm-go/lib/repository"
	"github.com/kopicee/hdm-go/lib/services"
	"golang.org/x/exp/slog"
)

func main() {
	hotelRepo := repository.NewHotelRepository()
	hotelsService := services.NewHotelsService(hotelRepo)

	mustIngest(hotelsService)

	api := api.NewAPI(hotelsService)
	mustListen(api, 3000)
}

func mustIngest(hotelsService services.HotelsService) {
	if err := hotelsService.Ingest(); err != nil {
		slog.Error("Failed to ingest supplier data, err: %v", err)
		panic(err)
	}
}

func mustListen(api api.API, port int) {
	if err := api.Listen(port); err != nil {
		slog.Error("Failed to start API, err: %v", err)
		panic(err)
	}
	slog.Info("Now listening on localhost:%d", port)
}
