package merge

import (
	"github.com/kopicee/hdm-go/lib/model"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slog"
)

func Merge(record *model.Hotel, updates []*model.Hotel) error {
	for _, next := range updates {
		if err := combine(record, next); err != nil {
			return err
		}
	}
	return nil
}

func combine(current *model.Hotel, next *model.Hotel) error {
	// Assume these don't change
	current.ID = next.ID
	current.DestinationID = next.DestinationID

	// For name, description and booking conditions, more verbose is better
	current.Name = chooseGreater(current.Name, next.Name, length)
	current.Description = chooseGreater(current.Description, next.Description, length)
	current.BookingConditions = chooseGreater(current.BookingConditions, next.BookingConditions, sliceLength[string])

	// For location, prefer more precision/verbosity
	current.Location.Latitude = chooseGreater(current.Location.Latitude, next.Location.Latitude, floatPrecision)
	current.Location.Longitude = chooseGreater(current.Location.Longitude, next.Location.Longitude, floatPrecision)
	current.Location.Address = chooseGreater(current.Location.Address, next.Location.Address, length)
	current.Location.City = chooseGreater(current.Location.City, next.Location.City, length)
	current.Location.Country = chooseGreater(current.Location.Country, next.Location.Country, length)

	// For amenities and images, join all data
	current.Amenities.General = append(current.Amenities.General, next.Amenities.General...)
	current.Amenities.Room = append(current.Amenities.Room, next.Amenities.Room...)
	current.Images.Site = append(current.Images.Site, next.Images.Site...)
	current.Images.Rooms = append(current.Images.Rooms, next.Images.Rooms...)
	current.Images.Amenities = append(current.Images.Amenities, next.Images.Amenities...)
	return nil
}

// chooseGreater comares the score of x and y, as calculated by the given scoreFunc, and returns
// the one which has a higher score. If they both share the same score, x is returned.
func chooseGreater[V any, S constraints.Ordered](x, y V, scoreFunc func(V) S, logger ...*slog.Logger) V {
	if scoreFunc(x) >= scoreFunc(y) {
		return x
	}
	return y
}
