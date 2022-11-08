package seq

// All determines whether all items of a sequence satisfy the given condition.
func (s seq[T]) All(p func(item T) bool) bool {
	it := s.iterator()
	for it.Next() {
		if !p(*it.Current()) {
			return false
		}
	}
	return true
}
