package seq

type cycledIterator[T any] struct {
	previous Iterator[T]
	buffer   []T
	index    int
	repeat   bool
}

func (i *cycledIterator[T]) Next() bool {
	if !i.repeat {
		ok := i.previous.Next()
		if ok {
			i.buffer = append(i.buffer, *i.previous.Current())
			return ok
		}
		i.repeat = true
	}
	i.index++
	if i.index == len(i.buffer) {
		i.index = 0
	}
	return i.index < len(i.buffer)
}

func (i *cycledIterator[T]) Current() *T {
	if !i.repeat {
		return i.previous.Current()
	}
	return &i.buffer[i.index]
}

// Cycled repeats a sequence infinitely.
func (s seq[T]) Cycled() Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &cycledIterator[T]{
				previous: s.Iterator(),
				index:    -1,
			}
		},
	}
}
