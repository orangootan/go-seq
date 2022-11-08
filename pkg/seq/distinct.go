package seq

type distinctIterator[T comparable] struct {
	previous Iterator[T]
	seen     map[T]bool
}

func (i *distinctIterator[T]) Next() bool {
	for i.previous.Next() {
		c := *i.previous.Current()
		if !i.seen[c] {
			i.seen[c] = true
			return true
		}
	}
	return false
}

func (i *distinctIterator[T]) Current() *T {
	return i.previous.Current()
}

// Distinct returns distinct elements from a sequence.
// This function works only with sequences of comparable items.
func Distinct[T comparable](s Seq[T]) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &distinctIterator[T]{
				previous: s.Iterator(),
				seen:     make(map[T]bool),
			}
		},
	}
}
