package seq

type cycleIterator[T any] struct {
	items []T
	index int
}

func (i *cycleIterator[T]) Next() bool {
	i.index++
	if i.index == len(i.items) {
		i.index = 0
	}
	return i.index < len(i.items)
}

func (i *cycleIterator[T]) Current() *T {
	return &i.items[i.index]
}

// Cycle returns an infinite sequence by cycling items from the given slice.
func Cycle[T any](items []T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &cycleIterator[T]{
				items: items,
				index: -1,
			}
		},
	}
}
