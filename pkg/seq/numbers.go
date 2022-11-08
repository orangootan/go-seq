package seq

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Real interface {
	Integer | Float
}

type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	Real | ~string
}

type Number interface {
	Real | Complex
}

// Sum returns sum of elements in a sequence.
// This function works with sequences of numbers.
func Sum[T Number](s Seq[T]) T {
	return s.Reduce(0, func(acc T, item T) T { return acc + item })
}

// Product returns product of elements in a sequence.
// This function works with sequences of numbers.
func Product[T Number](s Seq[T]) T {
	return s.Reduce(1, func(acc T, item T) T { return acc * item })
}

// Average returns average number of a sequence of numbers.
// This function works with sequences of integer and float numbers.
func Average[T Real](s Seq[T]) T {
	var count T
	sum := s.Reduce(0, func(acc T, item T) T {
		count++
		return acc + item
	})
	return sum / count
}

// Max returns the maximum element of a sequence and true if the sequence is not empty
// or undefined value and false otherwise.
// This function works with sequences of integer numbers, float numbers and strings.
func Max[T Ordered](s Seq[T]) (T, bool) {
	it := s.Iterator()
	var max T
	if !it.Next() {
		return max, false
	}
	max = *it.Current()
	for it.Next() {
		c := *it.Current()
		if c > max {
			max = c
		}
	}
	return max, true
}

// Min returns the minimum element of a sequence and true if the sequence is not empty
// or undefined value and false otherwise.
// This function works with sequences of integer numbers, float numbers and strings.
func Min[T Ordered](s Seq[T]) (T, bool) {
	it := s.Iterator()
	var min T
	if !it.Next() {
		return min, false
	}
	min = *it.Current()
	for it.Next() {
		c := *it.Current()
		if c < min {
			min = c
		}
	}
	return min, true
}
