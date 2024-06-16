package services

import (
	"github.com/kopicee/hdm-go/lib/model"
	"github.com/kopicee/hdm-go/lib/repository"
)

type HotelsService interface {
	Ingest() error
	Find(hotelIDs []string, destinationIDs []int) ([]*model.Hotel, error)
}

func NewHotelsService(repo repository.HotelsRepository) HotelsService {
	return hotelSvc{
		repo, // Implements Find()
		ingester{repo},
	}
}

type hotelSvc struct {
	repository.HotelsRepository
	ingester
}
