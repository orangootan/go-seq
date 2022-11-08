package seq

type chunkIterator[T any] struct {
	previous Iterator[T]
	size     int
	chunk    *[]T
}

func (i *chunkIterator[T]) Next() bool {
	var chunk []T
	for j := 0; j < i.size; j++ {
		if !i.previous.Next() {
			break
		}
		chunk = append(chunk, *i.previous.Current())
	}
	i.chunk = &chunk
	return len(chunk) > 0
}

func (i *chunkIterator[T]) Current() *[]T {
	return i.chunk
}

// Chunk splits a sequence into chunks of the given size and returns sequence of slices.
func Chunk[T any](size int, s Seq[T]) Seq[[]T] {
	return seq[[]T]{
		iterator: func() Iterator[[]T] {
			return &chunkIterator[T]{
				previous: s.Iterator(),
				size:     size,
			}
		},
	}
}
