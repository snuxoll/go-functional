package functional

import (
	"maps"
	"slices"
	"testing"
)

func TestSeq_Filter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	seq := OfSeq(slices.Values(input))

	result := seq.Filter(func(i int) bool {
		return i > 3
	}).Collect()

	expected := []int{4, 5}

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSeq_Filter2(t *testing.T) {
	input := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	seq := OfSeq2(maps.All(input))

	result := seq.Filter(func(key string, value int) bool {
		return value > 3
	}).Collect()

	expected := map[string]int{
		"four": 4,
		"five": 5,
	}

	if !maps.Equal(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
