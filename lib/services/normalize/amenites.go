package normalize

import (
	"strings"

	"github.com/fatih/camelcase"
	"github.com/kopicee/hdm-go/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

func normalizeAmenities(a model.Amenities) model.Amenities {
	return model.Amenities{
		General: functional.Map(a.General, normalizeAmenity),
		Room:    functional.Map(a.Room, normalizeAmenity),
	}
}

func normalizeAmenity(a model.Amenity) model.Amenity {
	normalized := strings.TrimSpace(string(a))
	if strings.EqualFold(normalized, "wifi") {
		return model.Amenity(normalized)
	}

	normalized = strings.Join(splitByCamelCase(normalized), " ")
	normalized = strings.ToLower(normalized)
	return model.Amenity(normalized)
}

func splitByCamelCase(s string) []string {
	return camelcase.Split(s)
}
