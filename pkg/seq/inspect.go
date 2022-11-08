package seq

type inspectIterator[T any] struct {
	previous Iterator[T]
	inspect  func(item T)
}

func (i *inspectIterator[T]) Next() bool {
	ok := i.previous.Next()
	if ok {
		i.inspect(*i.previous.Current())
	}
	return ok
}

func (i *inspectIterator[T]) Current() *T {
	return i.previous.Current()
}

// Inspect allows to observe items of a sequence by performing an action for each item.
func (s seq[T]) Inspect(do func(item T)) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &inspectIterator[T]{
				previous: s.iterator(),
				inspect:  do,
			}
		},
	}
}
