package seq

type mapIterator[T any, S any] struct {
	previous Iterator[T]
	selector func(item T) S
}

func (i *mapIterator[T, S]) Next() bool {
	return i.previous.Next()
}

func (i *mapIterator[T, S]) Current() *S {
	current := i.selector(*i.previous.Current())
	return &current
}

// Map projects items of a sequence using the given function.
func Map[T any, S any](s Seq[T], fn func(item T) S) Seq[S] {
	return seq[S]{
		iterator: func() Iterator[S] {
			return &mapIterator[T, S]{
				previous: s.Iterator(),
				selector: fn,
			}
		},
	}
}

// Select projects items of a sequence using the given function.
func Select[T any, S any](s Seq[T], fn func(item T) S) Seq[S] {
	return Map(s, fn)
}
