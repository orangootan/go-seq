package seq

type sliceReversedIterator[T any] struct {
	slice []T
	index int
}

func (i *sliceReversedIterator[T]) Next() bool {
	i.index--
	return i.index > -1
}

func (i *sliceReversedIterator[T]) Current() *T {
	return &i.slice[i.index]
}

// FromSliceReversed returns a sequence of items from the given slice in inverted order.
func FromSliceReversed[T any](slice []T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &sliceReversedIterator[T]{
				slice: slice,
				index: len(slice),
			}
		},
	}
}
