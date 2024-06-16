package suppliers

import (
	"context"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/kopicee/hdm-go/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

type patagoniaDTO struct {
	ID            string      `json:"ID"`
	DestinationID int         `json:"DestinationId"`
	Name          string      `json:"Name"`
	Lat           *Coordinate `json:"lat"`
	Lng           *Coordinate `json:"lng"`
	Address       string      `json:"Address"`
	Info          string      `json:"info"`
	Amenities     []string    `json:"amenities"`
	Images        struct {
		Rooms     []patagoniaImage `json:"room"`
		Amenities []patagoniaImage `json:"amenities"`
	} `json:"images"`
}

type patagoniaImage struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

func (*patagoniaDTO) adaptImage(p patagoniaImage) model.Image {
	return model.Image{Link: p.URL, Description: p.Description}
}

func (dto *patagoniaDTO) adapt() *model.Hotel {
	hotel := model.Hotel{}
	hotel.ID = dto.ID
	hotel.DestinationID = dto.DestinationID
	hotel.Name = dto.Name
	hotel.Description = dto.Info
	hotel.Location.Latitude = dto.Lat.Float64()
	hotel.Location.Longitude = dto.Lng.Float64()
	hotel.Location.Address = dto.Address
	hotel.Location.City = ""
	hotel.Location.Country = ""
	hotel.Amenities.General = []model.Amenity{}
	hotel.Amenities.Room = functional.Map(dto.Amenities, stringToAmenity)
	hotel.Images = model.Images{
		Rooms:     functional.Map(dto.Images.Rooms, dto.adaptImage),
		Site:      []model.Image{},
		Amenities: functional.Map(dto.Images.Amenities, dto.adaptImage),
	}
	hotel.BookingConditions = []string{}
	hotel.CreatedAt = time.Time{}
	hotel.UpdatedAt = time.Time{}
	return &hotel
}

type patagoniaSupplier struct {
	transport http.RoundTripper
}

func (s patagoniaSupplier) Name() string { return "patagonia" }

func (s patagoniaSupplier) Fetch(ctx context.Context) ([]*model.Hotel, error) {
	result := make([]*patagoniaDTO, 0)

	err := requests.URL("https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia").
		Method(http.MethodGet).
		ToJSON(&result).
		Transport(s.transport).
		Fetch(ctx)

	return adaptAll(result), err
}
