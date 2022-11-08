package seq

type skipIterator[T any] struct {
	previous        Iterator[T]
	toSkip, skipped int
}

func (i *skipIterator[T]) Next() bool {
	for i.skipped < i.toSkip && i.previous.Next() {
		i.skipped++
	}
	return i.previous.Next()
}

func (i *skipIterator[T]) Current() *T {
	return i.previous.Current()
}

// Skip bypasses the given number of items in a sequence and then returns the remaining items.
func (s seq[T]) Skip(n int) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &skipIterator[T]{
				previous: s.iterator(),
				toSkip:   n,
				skipped:  0,
			}
		},
	}
}
