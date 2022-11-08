package seq

type singleIterator[T any] struct {
	item T
	stop bool
}

func (i *singleIterator[T]) Next() bool {
	if i.stop {
		return false
	}
	i.stop = true
	return true
}

func (i *singleIterator[T]) Current() *T {
	return &i.item
}

// Single returns a sequence containing exactly one specified item.
func Single[T any](item T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &singleIterator[T]{
				item: item,
			}
		},
	}
}
