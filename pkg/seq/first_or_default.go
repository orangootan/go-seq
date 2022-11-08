package seq

// FirstOrDefault returns default value of type T if a sequence is empty
// or the first item otherwise.
func (s seq[T]) FirstOrDefault() T {
	it := s.iterator()
	if it.Next() {
		return *it.Current()
	}
	var t T
	return t
}
