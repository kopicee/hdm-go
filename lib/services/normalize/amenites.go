package normalize

import (
	"strings"

	"github.com/fatih/camelcase"
	"github.com/kopicee/hdm-go/lib/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

func normalizeAmenities(a model.Amenities) model.Amenities {
	return model.Amenities{
		General: functional.Map(a.General, normalizeAmenity),
		Room:    functional.Map(a.Room, normalizeAmenity),
	}
}

func normalizeAmenity(a model.Amenity) model.Amenity {
	s := string(a)
	s = strings.TrimSpace(s)

	if strings.EqualFold(s, "wifi") {
		return model.Amenity("wifi")
	}

	s = strings.ToLower(s)
	s = strings.Join(removeEmpty(camelcase.Split(s)), " ")
	s = strings.ToLower(s)
	return model.Amenity(s)
}
