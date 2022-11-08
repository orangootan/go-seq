package seq

// IsEmpty determines whether a sequence is empty.
func (s seq[T]) IsEmpty() bool {
	return !s.Iterator().Next()
}
