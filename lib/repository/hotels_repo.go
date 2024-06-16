package repository

import "github.com/kopicee/hdm-go/lib/model"

type HotelsRepository interface {
	Save(h *model.Hotel) error
	FindOne(id string) (*model.Hotel, error)
	Find(hotelIDs []string, destinationIDs []int) ([]*model.Hotel, error)
}

func NewHotelRepository() HotelsRepository {
	return &hotelsRepo{
		storage: make(map[string]*model.Hotel),
	}
}

type hotelsRepo struct {
	storage map[string]*model.Hotel
}

func (r *hotelsRepo) Save(h *model.Hotel) error {
	r.storage[h.ID] = h
	return nil
}

func (r *hotelsRepo) FindOne(id string) (*model.Hotel, error) {
	for _, h := range r.storage {
		if h.ID == id {
			return h, nil
		}
	}
	return nil, nil
}

func (r *hotelsRepo) Find(hotelIDs []string, destinationIDs []int) ([]*model.Hotel, error) {
	wantedHotels := toSet(hotelIDs)
	wantedDests := toSet(destinationIDs)
	noFilter := len(wantedHotels)+len(wantedDests) == 0

	ret := make([]*model.Hotel, 0)
	for _, h := range r.storage {
		if noFilter || wantedHotels.Contains(h.ID) || wantedDests.Contains(h.DestinationID) {
			ret = append(ret, h)
		}
	}
	return ret, nil
}

type Set[T comparable] map[T]bool

func (s Set[T]) Contains(value T) bool {
	_, found := s[value]
	return found
}

func toSet[T comparable](slice []T) Set[T] {
	ret := make(Set[T])
	for _, t := range slice {
		ret[t] = true
	}
	return ret
}
