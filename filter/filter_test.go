package filter

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	result := slices.Collect(Filter(slices.Values(input), func(i int) bool {
		return i > 3
	}))

	expected := []int{4, 5}

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
