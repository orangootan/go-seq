package seq

// Any determines whether any item of a sequence satisfies the given condition.
func (s seq[T]) Any(p func(item T) bool) bool {
	it := s.iterator()
	for it.Next() {
		if p(*it.Current()) {
			return true
		}
	}
	return false
}
