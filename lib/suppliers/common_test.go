package suppliers

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummy struct {
	Coord *Coordinate `json:"c"`
}

func Test_Coordinate_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		jsonDummy string
		expect    *Coordinate
	}{
		{`{"c":1.23}`, &Coordinate{value: 1.23, ok: true}},
		{`{"c":123}`, &Coordinate{value: 123.0, ok: true}},
		{`{"c":""}`, &Coordinate{value: 0, ok: false}},
		{`{"c":null}`, nil},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Unmarshal %s", tc.jsonDummy), func(t *testing.T) {
			dummy := dummy{}
			err := json.Unmarshal([]byte(tc.jsonDummy), &dummy)
			assert.NoError(t, err)

			actual := dummy.Coord
			if tc.expect == nil {
				assert.Nil(t, actual)
				return
			} else {
				assert.Equal(t, tc.expect.value, actual.value)
				assert.Equal(t, tc.expect.ok, actual.ok)
			}
		})
	}
}
