package suppliers

import (
	"context"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/kopicee/hdm-go/lib/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

type paperfliesDTO struct {
	HotelID       string `json:"hotel_id"`
	DestinationID int    `json:"destination_id"`
	HotelName     string `json:"hotel_name"`
	Location      struct {
		Address string `json:"address"`
		Country string `json:"country"`
	} `json:"location"`
	Details   string `json:"details"`
	Amenities struct {
		General []string `json:"general"`
		Room    []string `json:"room"`
	} `json:"amenities"`
	Images struct {
		Rooms []paperfliesImage `json:"rooms"`
		Site  []paperfliesImage `json:"site"`
	} `json:"images"`
	BookingConditions []string `json:"booking_conditions"`
}

type paperfliesImage struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

func (*paperfliesDTO) adaptImage(p paperfliesImage) model.Image {
	return model.Image{Link: p.Link, Description: p.Caption}
}

func (dto *paperfliesDTO) adapt() *model.Hotel {
	hotel := model.Hotel{}
	hotel.ID = dto.HotelID
	hotel.DestinationID = dto.DestinationID
	hotel.Name = dto.HotelName
	hotel.Description = dto.Details
	hotel.Location.Latitude = nil
	hotel.Location.Longitude = nil
	hotel.Location.Address = dto.Location.Address
	hotel.Location.City = ""
	hotel.Location.Country = dto.Location.Country
	hotel.Amenities.General = functional.Map(dto.Amenities.General, stringToAmenity)
	hotel.Amenities.Room = functional.Map(dto.Amenities.Room, stringToAmenity)
	hotel.Images = model.Images{
		Rooms:     functional.Map(dto.Images.Rooms, dto.adaptImage),
		Site:      functional.Map(dto.Images.Site, dto.adaptImage),
		Amenities: []model.Image{},
	}
	hotel.BookingConditions = dto.BookingConditions
	hotel.CreatedAt = time.Time{}
	hotel.UpdatedAt = time.Time{}
	return &hotel
}

type paperfliesSupplier struct {
	transport http.RoundTripper
}

func (s paperfliesSupplier) Name() string { return "paperflies" }

func (s paperfliesSupplier) Fetch(ctx context.Context) ([]*model.Hotel, error) {
	result := make([]*paperfliesDTO, 0)

	err := requests.URL("https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies").
		Method(http.MethodGet).
		ToJSON(&result).
		Transport(s.transport).
		Fetch(ctx)

	return adaptAll(result), err
}
