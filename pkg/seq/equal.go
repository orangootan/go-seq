package seq

// Equal determines whether two sequences have equal length and equal corresponding items.
// This function works only with sequences of comparable items.
func Equal[T comparable](s1 Seq[T], s2 Seq[T]) bool {
	it1 := s1.Iterator()
	it2 := s2.Iterator()
	for {
		ok1 := it1.Next()
		ok2 := it2.Next()
		if ok1 != ok2 {
			return false
		}
		if ok1 == false {
			return true
		}
		if *it1.Current() != *it2.Current() {
			return false
		}
	}
}
