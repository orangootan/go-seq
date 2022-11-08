package seq

// ToChan sends sequence items to an existing channel in current goroutine so this call is blocking.
// This function does not close the channel after the sequence is exhausted.
func (s seq[T]) ToChan(ch chan<- T) {
	it := s.iterator()
	for it.Next() {
		ch <- *it.Current()
	}
}
