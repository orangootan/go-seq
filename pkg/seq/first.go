package seq

// First returns undefined value and false if a sequence is empty
// or the first item and true otherwise.
func (s seq[T]) First() (T, bool) {
	it := s.iterator()
	if it.Next() {
		return *it.Current(), true
	}
	var t T
	return t, false
}
