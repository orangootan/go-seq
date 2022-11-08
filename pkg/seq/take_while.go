package seq

type takeWhileIterator[T any] struct {
	previous  Iterator[T]
	predicate func(item T) bool
	skip      bool
}

func (i *takeWhileIterator[T]) Next() bool {
	ok := !i.skip && i.previous.Next()
	if ok && !i.predicate(*i.previous.Current()) {
		i.skip = true
		return false
	}
	return ok
}

func (i *takeWhileIterator[T]) Current() *T {
	return i.previous.Current()
}

// TakeWhile returns items from a sequence as long as the given condition is true,
// and then skips the remaining items.
func (s seq[T]) TakeWhile(p func(item T) bool) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &takeWhileIterator[T]{
				previous:  s.iterator(),
				predicate: p,
			}
		},
	}
}
