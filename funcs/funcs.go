/*
Package funcs contains type-definitions and helpers for functions that this module accepts as parameters

# Filters

Filters are functions that take one (or more) arguments and return a boolean value. They are used to determine whether
an item should be included in a resulting sequence or not. Filters should return true if the item should be included,
or false if it should be included.

Example:

	myFilter := func(i int) bool {
		return i > 3
	}

	// myFilter(1) == false
	// myFilter(4) == true

# Combining filters

Filters can be combined to create more complex filters, and do so dynamically if desired. The [FilterFunc.Combine] and
[FilterPairFunc.Combine] methods can be used to return a new [FilterFunc] or [FilterPairFunc] that will return true if
any of the provided filters return true.

Example:

	firstFilter := func(i int) bool {
		return i > 3
	}

	secondFilter := func(i int) bool {
		return i == 1
	}

	combined := funcs.AsFilter(firstFilter).Combine(secondFilter)

	// combined(1) == true: secondFilter matched
	// combined(4) == true: firstFilter matched
	// combined(2) == false: neither filter matched

# Conversion helpers

Many function types in this package contain additional methods that can be called on them to manipulate and chain them
together, but to access them your function must be converted to the specific type first (Go will accept functions
matching the signature as arguments regardless of their declared type). Since this module makes heavy use of generics,
this package also provides helper functions that take advantage of Go's type inference to reduce (or eliminate) the need
for repeated definition of type parameters.

Example:

	myFilter := func(i int) bool {
		return i > 3
	}

	// Explicit conversion
	f := funcs.FilterFunc[int](myFilter)
	// f is now a funcs.FilterFunc[int]

	// Using helper
	f := funcs.AsFilter(myFilter)
	// f is now a funcs.FilterFunc[int]

List of conversion helpers:
  - [AsFilter]
  - [AsFilterPair]
*/
package funcs

// AsFilter converts a function conforming to the [FilterFunc] signature to a [FilterFunc] type.
func AsFilter[T any](fn FilterFunc[T]) FilterFunc[T] {
	return fn
}

// FilterFunc is a function that takes an item of type T and returns a boolean value, indicating whether the item should
// be included in the resulting sequence or not.
type FilterFunc[T any] func(T) bool

// Combine combines a FilterFunc with one or more additional FilterFuncs, returning a new FilterFunc that will
// return true if any of the provided functions return true.
//
// Example:
//
//	f1 := func(i int) bool {
//		return i > 3
//	}
//
//	f2 := func(i int) bool {
//		return i == 1
//	}
//
//	f := funcs.AsFilter(f1).Combine(f2)
//
//	// f(1) == true: f2 matched
//	// f(4) == true: f1 matched
//	// f(2) == false: neither filter matched
func (f FilterFunc[T]) Combine(filters ...FilterFunc[T]) FilterFunc[T] {
	return func(item T) bool {
		for _, filter := range filters {
			if filter(item) {
				return true
			}
		}
		return f(item)
	}
}

// Not returns a new FilterFunc that inverts the result of the original FilterFunc.
func (f FilterFunc[T]) Not() FilterFunc[T] {
	return func(item T) bool {
		return !f(item)
	}
}

// Unless returns a new FilterFunc that will return false if any of the provided filters return true, otherwise it will
// return the result of the original FilterFunc.
func (f FilterFunc[T]) Unless(filters ...FilterFunc[T]) FilterFunc[T] {
	return func(item T) bool {
		for _, filter := range filters {
			if filter(item) {
				return false
			}
		}
		return f(item)
	}
}

// AsFilterPair converts a function conforming to the [FilterPairFunc] signature to a [FilterPairFunc] type.
func AsFilterPair[K any, V any](fn FilterPairFunc[K, V]) FilterPairFunc[K, V] {
	return fn
}

// FilterPairFunc is a function that takes a key of type K and a value of type V and returns a boolean value, indicating
// whether the key-value pair should be included in the resulting sequence or not.
type FilterPairFunc[K any, V any] func(K, V) bool

// Combine combines a FilterPairFunc with one or more additional FilterPairFuncs, returning a new FilterPairFunc that
// will return true if any of the provided functions return true.
//
// Example:
//
//	f1 := func(key string, value int) bool {
//		return value > 3
//	}
//
//	f2 := func(key string, value int) bool {
//		return key == "one"
//	}
//
//	f := funcs.AsFilterPair(f1).Combine(f2)
//
//	// f("one", 1) == true: f2 matched
//	// f("four", 4) == true: f1 matched
//	// f("two", 2) == false: neither filter matched
func (f FilterPairFunc[K, V]) Combine(filters ...FilterPairFunc[K, V]) FilterPairFunc[K, V] {
	return func(key K, value V) bool {
		for _, filter := range filters {
			if filter(key, value) {
				return true
			}
		}
		return f(key, value)
	}
}

// Not returns a new FilterPairFunc that inverts the result of the original FilterPairFunc.
func (f FilterPairFunc[K, V]) Not() FilterPairFunc[K, V] {
	return func(key K, value V) bool {
		return !f(key, value)
	}
}

// Unless returns a new FilterPairFunc that will return false if any of the provided filters return true, otherwise it will
// return the result of the original FilterPairFunc.
func (f FilterPairFunc[K, V]) Unless(filters ...FilterPairFunc[K, V]) FilterPairFunc[K, V] {
	return func(key K, value V) bool {
		for _, filter := range filters {
			if filter(key, value) {
				return false
			}
		}
		return f(key, value)
	}
}

func AsMap[I any, O any](fn MapFunc[I, O]) MapFunc[I, O] {
	return fn
}

// MapFunc is a function that takes an item of type I and returns an item of type O, used to transform items in a
// sequence
type MapFunc[I any, O any] func(I) O

// Then chains additional MapFuncs to a MapFunc, returning a new MapFunc that will apply each function in sequence.
//
// Example:
//
//	f1 := func(i int) int {
//		return i * 2
//	}
//
//	f2 := func(i int) int {
//		return i + 1
//	}
//
//	f := funcs.AsMap(f1).Then(f2)
//
//	// f(1) == 3: f2(f1(1))
func (f MapFunc[I, O]) Then(fns ...MapFunc[O, O]) MapFunc[I, O] {
	return func(item I) O {
		result := f(item)
		for _, fn := range fns {
			result = fn(result)
		}
		return result
	}
}
