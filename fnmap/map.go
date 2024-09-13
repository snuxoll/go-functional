package fnmap

import (
	"go.snuxoll.com/functional/funcs"
	"iter"
)

func Map[I any, O any](in iter.Seq[I], fn funcs.MapFunc[I, O]) iter.Seq[O] {
	return func(yield func(out O) bool) {
		for val := range in {
			yield(fn(val))
		}
	}
}
