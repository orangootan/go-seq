package seq

// AsChan creates a channel and sends all sequence items into it in separate goroutine.
// After the sequence exhausted the channel is closed.
// You can specify the channel capacity with cap parameter.
func (s seq[T]) AsChan(cap int) <-chan T {
	it := s.iterator()
	ch := make(chan T, cap)
	go func() {
		for it.Next() {
			ch <- *it.Current()
		}
		close(ch)
	}()
	return ch
}
