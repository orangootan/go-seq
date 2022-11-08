package seq

// Iterator is the interface for iterator objects.
type Iterator[T any] interface {
	// Next advances iterator to the next item of the iterated sequence.
	// Returns true if successfully advanced or false if passed the end of the sequence.
	Next() bool
	// Current returns pointer to the current item of the iterated sequence.
	Current() *T
}

type seq[T any] struct {
	iterator func() Iterator[T]
}

// Iterator returns iterator object of a sequence.
func (s seq[T]) Iterator() Iterator[T] {
	return s.iterator()
}
