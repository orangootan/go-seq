package seq

type takeIterator[T any] struct {
	previous      Iterator[T]
	toTake, taken int
}

func (i *takeIterator[T]) Next() bool {
	if i.taken < i.toTake && i.previous.Next() {
		i.taken++
		return true
	}
	return false
}

func (i *takeIterator[T]) Current() *T {
	return i.previous.Current()
}

// Take returns the given number of items from the start of a sequence.
func (s seq[T]) Take(n int) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &takeIterator[T]{
				previous: s.iterator(),
				toTake:   n,
				taken:    0,
			}
		},
	}
}
