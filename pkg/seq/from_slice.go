package seq

type sliceIterator[T any] struct {
	slice []T
	index int
}

func (i *sliceIterator[T]) Next() bool {
	i.index++
	return i.index < len(i.slice)
}

func (i *sliceIterator[T]) Current() *T {
	return &i.slice[i.index]
}

// FromSlice returns a sequence of items from the given slice.
func FromSlice[T any](slice []T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &sliceIterator[T]{
				slice: slice,
				index: -1,
			}
		},
	}
}
