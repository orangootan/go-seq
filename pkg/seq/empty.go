package seq

type emptyIterator[T any] struct{}

func (i *emptyIterator[T]) Next() bool {
	return false
}

func (i *emptyIterator[T]) Current() *T {
	return nil
}

// Empty returns empty sequence.
func Empty[T any]() Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &emptyIterator[T]{}
		},
	}
}
