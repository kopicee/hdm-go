package normalize

import (
	"github.com/kopicee/hdm-go/functional"
	"github.com/kopicee/hdm-go/lib/model"
)

func normalizeImages(i model.Images) model.Images {
	linkFunc := func(img model.Image) string { return img.Link }

	return model.Images{
		Rooms:     functional.RemoveDuplicatesBy(i.Rooms, linkFunc),
		Site:      functional.RemoveDuplicatesBy(i.Site, linkFunc),
		Amenities: functional.RemoveDuplicatesBy(i.Amenities, linkFunc),
	}
}
