package seq

// ForEnumerated is similar to ForEach but passes item index as first parameter of action function.
func (s seq[T]) ForEnumerated(do func(i int, item T)) {
	it := s.iterator()
	i := -1
	for it.Next() {
		i++
		do(i, *it.Current())
	}
}
