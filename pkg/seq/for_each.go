package seq

// ForEach performs an action for each item in a sequence.
func (s seq[T]) ForEach(do func(item T)) {
	it := s.iterator()
	for it.Next() {
		do(*it.Current())
	}
}
