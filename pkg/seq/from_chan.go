package seq

type chanIterator[T any] struct {
	ch      chan T
	current *T
}

func (i *chanIterator[T]) Next() bool {
	item, ok := <-i.ch
	if !ok {
		return false
	}
	i.current = &item
	return true
}

func (i *chanIterator[T]) Current() *T {
	return i.current
}

// FromChan returns a sequence of items received from a channel until the channel is closed.
func FromChan[T any](ch chan T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &chanIterator[T]{
				ch: ch,
			}
		},
	}
}
