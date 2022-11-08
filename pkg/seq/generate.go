package seq

type generateIterator[T any] struct {
	fn func(i int) T
	i  int
}

func (i *generateIterator[T]) Next() bool {
	i.i++
	return true
}

func (i *generateIterator[T]) Current() *T {
	item := i.fn(i.i)
	return &item
}

// Generate returns an infinite sequence applying a mapping function to item indices.
func Generate[T any](fn func(i int) T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &generateIterator[T]{
				fn: fn,
				i:  -1,
			}
		},
	}
}
