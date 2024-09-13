package combine

import (
	"go.snuxoll.com/functional/filter"
	"maps"
	"slices"
	"testing"
)

func TestFilterFunc(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	f1 := func(i int) bool {
		return i > 3
	}

	f2 := func(i int) bool {
		return i == 1
	}

	f := FilterFunc(f1, f2)

	result := slices.Collect(filter.Filter(slices.Values(input), f))

	expected := []int{1, 4, 5}

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFilterPairFunc(t *testing.T) {
	input := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	f1 := func(key string, value int) bool {
		return value > 3
	}

	f2 := func(key string, value int) bool {
		return key == "one"
	}

	f := FilterPairFunc(f1, f2)

	result := maps.Collect(filter.FilterPair(maps.All(input), f))

	expected := map[string]int{
		"one":  1,
		"four": 4,
		"five": 5,
	}

	if !maps.Equal(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
