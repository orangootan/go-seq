package seq

// FlatMap maps each sequence item to a sequence and flattens results.
func FlatMap[T any, R any](s Seq[T], fn func(item T) Seq[R]) Seq[R] {
	return Flatten(Map(s, fn))
}
