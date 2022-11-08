package seq

type iterateIterator[T any] struct {
	fn      func(T) T
	seed    T
	current *T
}

func (i *iterateIterator[T]) Next() bool {
	if i.current == nil {
		i.current = &i.seed
	} else {
		c := i.fn(*i.current)
		i.current = &c
	}
	return true
}

func (i *iterateIterator[T]) Current() *T {
	return i.current
}

// Iterate returns an infinite sequence consisting of seed, fn(seed), fn(fn(seed)), etc.
func Iterate[T any](seed T, fn func(T) T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &iterateIterator[T]{
				fn:   fn,
				seed: seed,
			}
		},
	}
}
