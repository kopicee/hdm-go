package suppliers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/kopicee/hdm-go/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

type acmeDTO struct {
	ID            string      `json:"ID"`
	DestinationID int         `json:"DestinationId"`
	Name          string      `json:"Name"`
	Latitude      *Coordinate `json:"Latitude"`
	Longitude     *Coordinate `json:"Longitude"`
	Address       string      `json:"Address"`
	City          string      `json:"City"`
	Country       string      `json:"Country"`
	PostalCode    string      `json:"PostalCode"`
	Description   string      `json:"Description"`
	Facilities    []string    `json:"Facilities"`
}

func (dto *acmeDTO) adapt() *model.Hotel {
	hotel := model.Hotel{}
	hotel.ID = dto.ID
	hotel.DestinationID = dto.DestinationID
	hotel.Name = dto.Name
	hotel.Description = dto.Description
	hotel.Location.Latitude = dto.Latitude.Float64()
	hotel.Location.Longitude = dto.Longitude.Float64()
	hotel.Location.Address = dto.joinAddress(dto.Address, dto.PostalCode)
	hotel.Location.City = dto.City
	hotel.Location.Country = dto.Country
	hotel.Amenities.General = functional.Map(dto.Facilities, stringToAmenity)
	hotel.Amenities.Room = []model.Amenity{}
	hotel.Images = model.Images{
		Rooms:     []model.Image{},
		Site:      []model.Image{},
		Amenities: []model.Image{},
	}
	hotel.BookingConditions = []string{}
	hotel.CreatedAt = time.Time{}
	hotel.UpdatedAt = time.Time{}
	return &hotel
}

func (dto *acmeDTO) joinAddress(address, postcode string) string {
	if strings.Contains(address, postcode) {
		return address
	}
	if postcode == "" {
		return address
	}
	return fmt.Sprintf("%s, %s", address, postcode)
}

type acmeSupplier struct {
	transport http.RoundTripper
}

func (s acmeSupplier) Name() string { return "acme" }

func (s acmeSupplier) Fetch(ctx context.Context) ([]*model.Hotel, error) {
	result := make([]*acmeDTO, 0)

	err := requests.URL("https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme").
		Method(http.MethodGet).
		ToJSON(&result).
		Transport(s.transport).
		Fetch(ctx)

	return adaptAll(result), err
}
