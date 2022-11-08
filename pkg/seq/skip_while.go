package seq

type skipWhileIterator[T any] struct {
	previous  Iterator[T]
	predicate func(item T) bool
	take      bool
}

func (i *skipWhileIterator[T]) Next() bool {
	if i.take {
		return i.previous.Next()
	}
	for i.previous.Next() {
		if !i.predicate(*i.previous.Current()) {
			i.take = true
			return true
		}
	}
	return false
}

func (i *skipWhileIterator[T]) Current() *T {
	return i.previous.Current()
}

// SkipWhile bypasses items in a sequence as long as the given condition is true
// and then returns the remaining items.
func (s seq[T]) SkipWhile(p func(item T) bool) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &skipWhileIterator[T]{
				previous:  s.iterator(),
				predicate: p,
			}
		},
	}
}
