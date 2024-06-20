package suppliers

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/kopicee/hdm-go/lib/model"
)

type adapter interface {
	adapt() *model.Hotel
}

func adaptAll[T adapter](adapters []T) []*model.Hotel {
	result := make([]*model.Hotel, len(adapters))
	for i, dto := range adapters {
		result[i] = dto.adapt()
	}
	return result
}

func stringToAmenity(s string) model.Amenity {
	return model.Amenity(s)
}

type Coordinate struct {
	value float64
	ok    bool
}

func (c *Coordinate) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		c.ok = false
		return nil
	}

	if data[0] == '"' {
		value, err := strconv.ParseFloat(string(data), 64)
		if err != nil || math.IsInf(value, 0) || math.IsNaN(value) {
			c.ok = false
			return nil
		}
		c.value = value
	} else {
		var value float64
		if err := json.Unmarshal(data, &value); err != nil {
			c.ok = false
			return nil
		}
		c.value = value
	}

	c.ok = true
	return nil
}

func (c *Coordinate) Float64() *float64 {
	if c == nil || !c.ok {
		return nil
	}
	return &c.value
}
