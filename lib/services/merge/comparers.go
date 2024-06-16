package merge

import (
	"fmt"
	"strings"
)

func length(s string) int { return len(s) }

func sliceLength[T any](slice []T) int { return len(slice) }

func floatPrecision(f *float64) int {
	if f == nil {
		return 0
	}
	parts := strings.Split(fmt.Sprint(*f), ".")
	if len(parts) < 2 {
		return 0
	}
	return len(parts[1])
}
