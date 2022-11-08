package seq

// AsSlice collects all sequence items into newly created slice.
func (s seq[T]) AsSlice() []T {
	it := s.iterator()
	var acc []T
	for it.Next() {
		acc = append(acc, *it.Current())
	}
	return acc
}
