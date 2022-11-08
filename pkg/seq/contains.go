package seq

// Contains determines whether a sequence contains the given item.
// This function works only with sequences of comparable items.
func Contains[T comparable](item T, s Seq[T]) bool {
	it := s.Iterator()
	for it.Next() {
		if *it.Current() == item {
			return true
		}
	}
	return false
}
