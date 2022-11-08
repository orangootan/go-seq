package seq

// GroupBy returns a map grouping sequence items by key produced by the given function.
// Key type must be comparable.
func GroupBy[T any, K comparable](s Seq[T], key func(item T) K) map[K][]T {
	m := make(map[K][]T)
	it := s.Iterator()
	for it.Next() {
		item := *it.Current()
		k := key(item)
		m[k] = append(m[k], item)
	}
	return m
}
