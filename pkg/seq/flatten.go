package seq

type flattenIterator[T any] struct {
	previous Iterator[Seq[T]]
	current  Iterator[T]
}

func (i *flattenIterator[T]) Next() bool {
	ok := i.current != nil && i.current.Next()
	if !ok {
		for i.previous.Next() {
			i.current = (*i.previous.Current()).Iterator()
			if i.current.Next() {
				return true
			}
		}
	}
	return ok
}

func (i *flattenIterator[T]) Current() *T {
	return i.current.Current()
}

// Flatten combines a sequence of item sequences into flat sequence of items.
func Flatten[T any](s Seq[Seq[T]]) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &flattenIterator[T]{
				previous: s.Iterator(),
			}
		},
	}
}
