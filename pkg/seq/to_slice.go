package seq

// ToSlice appends items of a sequence to an existing slice given by reference.
func (s seq[T]) ToSlice(acc *[]T) {
	it := s.iterator()
	for it.Next() {
		*acc = append(*acc, *it.Current())
	}
}
