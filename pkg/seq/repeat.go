package seq

type repeatIterator[T any] struct {
	item T
}

func (i *repeatIterator[T]) Next() bool {
	return true
}

func (i *repeatIterator[T]) Current() *T {
	return &i.item
}

// Repeat returns an infinite sequence repeating the given item.
func Repeat[T any](item T) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &repeatIterator[T]{
				item: item,
			}
		},
	}
}
