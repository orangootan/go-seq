package seq

type filterIterator[T any] struct {
	previous  Iterator[T]
	predicate func(item T) bool
}

func (i *filterIterator[T]) Next() bool {
	for i.previous.Next() {
		if i.predicate(*i.previous.Current()) {
			return true
		}
	}
	return false
}

func (i *filterIterator[T]) Current() *T {
	return i.previous.Current()
}

// Filter filters a sequence of items based on the given predicate.
func (s seq[T]) Filter(p func(item T) bool) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &filterIterator[T]{
				previous:  s.iterator(),
				predicate: p,
			}
		},
	}
}

// Where filters a sequence of items based on the given predicate.
func (s seq[T]) Where(p func(item T) bool) Seq[T] {
	return s.Filter(p)
}
