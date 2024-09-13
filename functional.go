package functional

import (
	"go.snuxoll.com/functional/filter"
	"go.snuxoll.com/functional/fnmap"
	"go.snuxoll.com/functional/funcs"
	"iter"
	"maps"
	"slices"
)

/*
OfSeq casts an [iter.Seq] to a [Seq] for use with the functional package.
*/
func OfSeq[T any](seq iter.Seq[T]) Seq[T] {
	return Seq[T](seq)
}

/*
OfSeq2 casts an [iter.Seq2] to a [Seq2] for use with the functional package.
*/
func OfSeq2[K comparable, V any](seq iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](seq)
}

/*
Seq is an alias for [iter.Seq] that provides functional easy access to other functions of the [functional] package
as methods that can be easily chained together.
*/
type Seq[T any] iter.Seq[T]

/*
Seq2 is an alias for [iter.Seq2] that provides functional easy access to other functions of the [functional] package
as methods that can be easily chained together.
*/
type Seq2[K comparable, V any] iter.Seq2[K, V]

/*
Map applies the provided function to each element of the input sequence, returning a new sequence containing the
results.
*/
func Map[I any, O any](seq Seq[I], f funcs.MapFunc[I, O]) Seq[O] {
	return Seq[O](fnmap.Map(iter.Seq[I](seq), f))
}

/*
Filter filters the input sequence based on the provided filter function, returning a new sequence containing only
the elements that pass the filter.
*/
func (s Seq[T]) Filter(f funcs.FilterFunc[T]) Seq[T] {
	return Seq[T](filter.Filter(iter.Seq[T](s), f))
}

/*
Collect returns the elements of the sequence as a slice.
*/
func (s Seq[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](s))
}

/*
Filter filters the input sequence based on the provided filter function, returning a new sequence containing only
the elements that pass the filter.
*/
func (s Seq2[K, V]) Filter(f funcs.FilterPairFunc[K, V]) Seq2[K, V] {
	return Seq2[K, V](filter.FilterPair(iter.Seq2[K, V](s), f))
}

/*
Collect returns the elements of the sequence as a map.
*/
func (s Seq2[K, V]) Collect() map[K]V {
	return maps.Collect(iter.Seq2[K, V](s))
}
