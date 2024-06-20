package normalize

import (
	"strings"

	"github.com/kopicee/hdm-go/lib/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

func Normalize(h *model.Hotel) error {
	h.Name = strings.TrimSpace(h.Name)
	h.Description = strings.TrimSpace(h.Description)
	h.Location = normalizeLocation(h.Location)
	h.Amenities = normalizeAmenities(h.Amenities)
	h.Images = normalizeImages(h.Images)

	return nil
}

func removeEmpty(elems []string) []string {
	return functional.Filter(
		elems,
		func(s string) bool { return len(strings.TrimSpace(s)) != 0 },
	)
}
