package seq

// Fold applies an accumulator function over a sequence given a specified initial value.
// Accumulator and item types can be different.
func Fold[T any, R any](s Seq[T], acc R, fn func(R, T) R) R {
	it := s.Iterator()
	for it.Next() {
		acc = fn(acc, *it.Current())
	}
	return acc
}
