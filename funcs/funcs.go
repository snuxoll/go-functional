package funcs

type FilterFunc[T any] func(T) bool

type FilterPairFunc[K any, V any] func(K, V) bool

type MapFunc[I any, O any] func(I) O
