package seq

type interleaveIterator[T any] struct {
	previous1 Iterator[T]
	previous2 Iterator[T]
	second    bool
}

func (i *interleaveIterator[T]) Next() bool {
	if i.second {
		ok := i.previous2.Next()
		if ok {
			i.second = false
		}
		return ok
	}
	ok := i.previous1.Next()
	if ok {
		i.second = true
	}
	return ok
}

func (i *interleaveIterator[T]) Current() *T {
	if i.second {
		return i.previous1.Current()
	}
	return i.previous2.Current()
}

// Interleave returns a sequence that interleaves items from two sequences.
func (s seq[T]) Interleave(s2 Seq[T]) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &interleaveIterator[T]{
				previous1: s.iterator(),
				previous2: s2.Iterator(),
			}
		},
	}
}
