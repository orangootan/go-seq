package seq

// Partition returns two slices. The first one contains items satisfying the given condition,
// the second one contains the rest of them.
func (s seq[T]) Partition(p func(item T) bool) ([]T, []T) {
	it := s.iterator()
	var t, f []T
	for it.Next() {
		c := *it.Current()
		if p(c) {
			t = append(t, c)
		} else {
			f = append(f, c)
		}
	}
	return t, f
}
