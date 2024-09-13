package combine

import (
	"go.snuxoll.com/functional/funcs"
	"slices"
)

/*
FilterFunc combines multiple filter functions into a single filter function, which will return true if any of the
provided functions returns true (implementing a logical OR).

Example:

	input := []int{1, 2, 3, 4, 5}

	f1 := func(i int) bool {
		return i > 3
	}

	f2 := func(i int) bool {
		return i == 1
	}

	f := combine.FilterFunc(f1, f2)

	result := slices.Collect(filter.Filter(slices.Values(input), f))
	// result == []int{1, 4, 5}
*/
func FilterFunc[T any](filters ...funcs.FilterFunc[T]) funcs.FilterFunc[T] {
	return func(item T) bool {
		return slices.ContainsFunc(filters, func(filter funcs.FilterFunc[T]) bool {
			return filter(item)
		})
	}
}

/*
FilterPairFunc combines multiple filter functions into a single filter function, which will return true if any of the
provided functions returns true (implementing a logical OR).

Example:

	input := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
	}

	f1 := func(key string, value int) bool {
		return value > 3
	}

	f2 := func(key string, value int) bool {
		return key == "one"
	}

	f := combine.FilterPairFunc(f1, f2)

	result := slices.Collect(filter.FilterPair(slices.Pairs(input), f))
	// result == map[string]int{"one": 1, "four": 4, "five": 5}
*/
func FilterPairFunc[K any, V any](filters ...funcs.FilterPairFunc[K, V]) funcs.FilterPairFunc[K, V] {
	return func(key K, value V) bool {
		return slices.ContainsFunc(filters, func(filter funcs.FilterPairFunc[K, V]) bool {
			return filter(key, value)
		})
	}
}
