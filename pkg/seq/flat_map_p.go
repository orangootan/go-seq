package seq

// FlatMapP maps each sequence item to a sequence and flattens results.
// This is parallel version of FlatMap using specified number of goroutines to process items in parallel.
func FlatMapP[T any, R any](par int, s Seq[T], fn func(item T) Seq[R]) Seq[R] {
	return Flatten(MapP(par, s, fn))
}
