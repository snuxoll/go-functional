package functional

import (
	"go.snuxoll.com/functional/filter"
	"go.snuxoll.com/functional/funcs"
	"iter"
	"maps"
	"slices"
)

// AsFilter converts a function conforming to the [FilterFunc] signature to a [FilterFunc] type.
func AsFilter[T any](f funcs.FilterFunc[T]) funcs.FilterFunc[T] {
	return f
}

// AsFilterPair converts a function conforming to the [FilterPairFunc] signature to a [FilterPairFunc] type.
func AsFilterPair[K any, V any](f funcs.FilterPairFunc[K, V]) funcs.FilterPairFunc[K, V] {
	return f
}

/*
Pair is a simple interface for key-value pairs. It is used to represent key-value pairs in a way that can be used with
functions expecting the [Seq] type.
*/
type Pair[K any, V any] interface {
	Key() K
	Value() V
}

type pair[K any, V any] struct {
	key   K
	value V
}

type inputSeq[T any] interface {
	iter.Seq[T] | Seq[T]
}

func (p pair[K, V]) Key() K {
	return p.key
}

func (p pair[K, V]) Value() V {
	return p.value
}

func PairOf[K any, V any](key K, value V) Pair[K, V] {
	return pair[K, V]{key, value}
}

// OfSlice converts a slice to a [Seq] for use with the functional package.
func OfSlice[T any](slice []T) Seq[T] {
	return Seq[T](slices.Values(slice))
}

// OfSeq casts an [iter.Seq] to a [Seq] for use with the functional package.
func OfSeq[T any](seq iter.Seq[T]) Seq[T] {
	return Seq[T](seq)
}

// OfSeq2 casts an [iter.Seq2] to a [Seq2] for use with the functional package.
func OfSeq2[K comparable, V any](seq iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](seq)
}

// OfPairSeq casts a [Seq] of key-value pairs to a [Seq2] for use with the functional package.
func OfPairSeq[K comparable, V any, S iter.Seq[Pair[K, V]] | Seq[Pair[K, V]]](seq S) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for p := range seq {
			if !yield(p.Key(), p.Value()) {
				return
			}
		}
	}
}

// Seq is a new type for [iter.Seq] that provides functional easy access to other functions of the [functional] package
// as methods that can be easily chained together.
type Seq[T any] iter.Seq[T]

// Seq2 is a new type for [iter.Seq2] that provides functional easy access to other functions of the [functional] package
// as methods that can be easily chained together.
type Seq2[K comparable, V any] iter.Seq2[K, V]

// Std converts the [Seq] to an [iter.Seq] for use with other packages.
func (s Seq[T]) Std() iter.Seq[T] {
	return iter.Seq[T](s)
}

// Filter filters the input sequence based on the provided filter function, returning a new sequence containing only
// the elements that pass the filter.
func (s Seq[T]) Filter(f funcs.FilterFunc[T]) Seq[T] {
	return Seq[T](filter.Seq(iter.Seq[T](s), f))
}

// Collect returns the elements of the sequence as a slice.
func (s Seq[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](s))
}

// Filter filters the input sequence based on the provided filter function, returning a new sequence containing only
// the elements that pass the filter.
func (s Seq2[K, V]) Filter(f funcs.FilterPairFunc[K, V]) Seq2[K, V] {
	return Seq2[K, V](filter.Seq2(iter.Seq2[K, V](s), f))
}

// Collect returns the elements of the sequence as a map.
func (s Seq2[K, V]) Collect() map[K]V {
	return maps.Collect(iter.Seq2[K, V](s))
}

// Std converts the [Seq2] to an [iter.Seq2] for use with other packages.
func (s Seq2[K, V]) Std() iter.Seq2[K, V] {
	return iter.Seq2[K, V](s)
}

// AsSeq converts a Seq2 to a Seq of key-value pairs implementing the Pair interface
func (s Seq2[K, V]) AsSeq() Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for key, value := range s {
			if !yield(PairOf(key, value)) {
				return
			}
		}
	}
}

// Map applies the provided function to each element of the input sequence, returning a new sequence containing the
// resulting values.
func Map[I any, O any](in Seq[I], fn funcs.MapFunc[I, O]) Seq[O] {
	return func(yield func(out O) bool) {
		for val := range in {
			yield(fn(val))
		}
	}
}

type concatSeq[T any] interface {
	sequences() []Seq[T]
}

type concatInput[T any] interface {
	[]Seq[T] | Seq[Seq[T]]
}

// Concat concatenates the provided sequences into a single sequence. This function does return elements in order,
// starting with the first element of the first sequence and ending with the last element of the last sequence. This
// does not guarantee that the underlying sequences give a predictable order.
func Concat[T any](seqs Seq[Seq[T]]) Seq[T] {

	return func(yield func(T) bool) {
		for seq := range seqs {
			for val := range seq {
				if !yield(val) {
					return
				}
			}
		}
	}
}
