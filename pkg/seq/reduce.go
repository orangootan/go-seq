package seq

// Reduce applies an accumulator function over a sequence given a specified initial value.
func (s seq[T]) Reduce(acc T, fn func(acc T, item T) T) T {
	it := s.iterator()
	for it.Next() {
		acc = fn(acc, *it.Current())
	}
	return acc
}
