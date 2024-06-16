package model

import (
	"time"
)

type Hotel struct {
	ID                string    `json:"id"`
	DestinationID     int       `json:"destination_id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	BookingConditions []string  `json:"booking_conditions"`
	Location          Location  `json:"location"`
	Amenities         Amenities `json:"amenities"`
	Images            Images    `json:"images"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Location struct {
	Latitude  *float64 `json:"lat"`
	Longitude *float64 `json:"lng"`
	Address   string   `json:"address"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
}

type Amenities struct {
	General []Amenity `json:"general"`
	Room    []Amenity `json:"room"`
}

type Amenity string

type Images struct {
	Rooms     []Image `json:"rooms"`
	Site      []Image `json:"site"`
	Amenities []Image `json:"amenities"`
}

type Image struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}
