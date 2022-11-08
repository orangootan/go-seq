package seq

type reverseIterator[T any] struct {
	previous  Iterator[T]
	buffer    []T
	collected bool
	index     int
}

func (i *reverseIterator[T]) Next() bool {
	if !i.collected {
		for i.previous.Next() {
			i.buffer = append(i.buffer, *i.previous.Current())
		}
		i.collected = true
		i.index = len(i.buffer)
	}
	i.index--
	return i.index > -1
}

func (i *reverseIterator[T]) Current() *T {
	return &i.buffer[i.index]
}

// Reverse inverts the order of items in a sequence.
func (s seq[T]) Reverse() Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &reverseIterator[T]{
				previous: s.iterator(),
			}
		},
	}
}
