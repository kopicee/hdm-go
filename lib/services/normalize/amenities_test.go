package normalize

import (
	"testing"

	"github.com/kopicee/hdm-go/lib/model"
	"github.com/stretchr/testify/assert"
)

func Test_normalizeAmenity(t *testing.T) {
	testCases := []struct {
		input, output string
	}{
		{input: "WiFi", output: "wifi"},
		{input: " Indoor Pool", output: "indoor pool"},
	}
	for _, tc := range testCases {
		t.Run("normalize "+tc.input, func(t *testing.T) {
			expect := model.Amenity(tc.output)
			actual := normalizeAmenity(model.Amenity(tc.input))

			assert.Equal(t, expect, actual)
		})
	}
}
