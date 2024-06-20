package normalize

import (
	"strings"

	"github.com/biter777/countries"
	"github.com/kopicee/hdm-go/lib/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

func normalizeLocation(loc model.Location) model.Location {
	lat, lng := normalizeCoords(loc.Latitude, loc.Longitude)
	return model.Location{
		Latitude:  lat,
		Longitude: lng,
		Address:   normalizeAddress(loc.Address),
		City:      normalizeCity(loc.City),
		Country:   normalizeCountry(loc.Country),
	}
}

func normalizeCoords(lat, lng *float64) (*float64, *float64) {
	invalid := func(c *float64) bool {
		return c == nil || *c < -180 || *c > 180
	}

	if invalid(lat) || invalid(lng) {
		return nil, nil
	}

	return lat, lng
}

func normalizeAddress(s string) string {
	parts := strings.Split(s, ",")
	parts = removeEmpty(functional.Map(parts, strings.TrimSpace))
	return strings.Join(parts, ", ")
}

func normalizeCity(s string) string {
	return strings.TrimSpace(s)
}

func normalizeCountry(s string) string {
	country := countries.ByName(s)
	if country == countries.Unknown {
		return ""
	}
	return country.Info().Name
}
