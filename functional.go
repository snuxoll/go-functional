package functional

import (
	"go.snuxoll.com/functional/filter"
	"go.snuxoll.com/functional/funcs"
	"iter"
	"maps"
	"slices"
)

func OfSeq[T any](seq iter.Seq[T]) Seq[T] {
	return Seq[T](seq)
}

func OfSeq2[K comparable, V any](seq iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](seq)
}

type Seq[T any] iter.Seq[T]

type Seq2[K comparable, V any] iter.Seq2[K, V]

func (s Seq[T]) Filter(f funcs.FilterFunc[T]) Seq[T] {
	return Seq[T](filter.Filter(iter.Seq[T](s), f))
}

func (s Seq[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](s))
}

func (s Seq2[K, V]) Filter(f funcs.FilterPairFunc[K, V]) Seq2[K, V] {
	return Seq2[K, V](filter.FilterPair(iter.Seq2[K, V](s), f))
}

func (s Seq2[K, V]) Collect() map[K]V {
	return maps.Collect(iter.Seq2[K, V](s))
}
