package seq

// Count counts number of items in a sequence.
// You can specify an integer type for result.
func Count[T any, C Integer](s Seq[T]) C {
	it := s.Iterator()
	var count C
	for it.Next() {
		count++
	}
	return count
}
