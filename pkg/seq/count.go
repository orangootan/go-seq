package seq

// Count counts number of items in a sequence.
func (s seq[T]) Count() int {
	it := s.iterator()
	var count int
	for it.Next() {
		count++
	}
	return count
}
