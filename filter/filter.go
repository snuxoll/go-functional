// Package filter provides support for filtering [iter.Seq]
package filter

import (
	"go.snuxoll.com/functional/funcs"
	"iter"
)

func Filter[T any, S iter.Seq[T]](seq S, filter funcs.FilterFunc[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range seq {
			if !filter(item) {
				continue
			}
			if !yield(item) {
				return
			}
		}
	}
}

//goland:noinspection GoNameStartsWithPackageName
func FilterPair[K any, V any, S iter.Seq2[K, V]](seq S, filter funcs.FilterPairFunc[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range seq {
			if !filter(key, value) {
				continue
			}
			if !yield(key, value) {
				return
			}
		}
	}
}
