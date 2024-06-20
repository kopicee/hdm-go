package normalize

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type coords struct{ lat, lng *float64 }

func (c coords) toSlice() []*float64 {
	return []*float64{c.lat, c.lng}
}

func Test_normalizeCoords(t *testing.T) {
	ok := 1.0
	over := 180.1
	under := -180.1
	invalidCoords := coords{nil, nil}
	testCases := []struct {
		input, expect coords
	}{
		{input: coords{&ok, &ok}, expect: coords{&ok, &ok}},
		{input: coords{&ok, nil}, expect: invalidCoords},
		{input: coords{nil, &ok}, expect: invalidCoords},
		{input: coords{nil, nil}, expect: invalidCoords},
		{input: coords{nil, nil}, expect: invalidCoords},
		{input: coords{&over, nil}, expect: invalidCoords},
		{input: coords{&under, nil}, expect: invalidCoords},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("normalize (%v, %v)", tc.input.lat, tc.input.lng), func(t *testing.T) {
			actualLat, actualLng := normalizeCoords(tc.input.lat, tc.input.lng)
			actual := coords{actualLat, actualLng}

			assert.Equal(t, tc.expect.toSlice(), actual.toSlice())
		})
	}
}

func Test_normalizeAddress(t *testing.T) {
	testCases := []struct {
		input, expect string
	}{
		{
			input:  " 5 Marine Way, ",
			expect: "5 Marine Way",
		},
		{
			input:  " 5 Marine Way   ",
			expect: "5 Marine Way",
		},
	}
	for _, tc := range testCases {
		t.Run("normalize "+tc.input, func(t *testing.T) {
			actual := normalizeAddress(tc.input)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

func Test_normalizeCountry(t *testing.T) {
	testCases := []struct {
		input, expect string
	}{
		{
			input:  "SG",
			expect: "Singapore",
		},
		{
			input:  "Singapore",
			expect: "Singapore",
		},
		{
			input:  "",
			expect: "",
		},
	}
	for _, tc := range testCases {
		t.Run("normalize "+tc.input, func(t *testing.T) {
			actual := normalizeCountry(tc.input)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
