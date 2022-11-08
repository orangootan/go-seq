package seq

type concatIterator[T any] struct {
	previous1 Iterator[T]
	previous2 Iterator[T]
	second    bool
}

func (i *concatIterator[T]) Next() bool {
	if i.second {
		return i.previous2.Next()
	}
	ok := i.previous1.Next()
	if !ok {
		i.second = true
		ok = i.previous2.Next()
	}
	return ok
}

func (i *concatIterator[T]) Current() *T {
	if i.second {
		return i.previous2.Current()
	}
	return i.previous1.Current()
}

// Concat concatenates two sequences.
func (s seq[T]) Concat(s2 Seq[T]) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &concatIterator[T]{
				previous1: s.iterator(),
				previous2: s2.Iterator(),
			}
		},
	}
}
