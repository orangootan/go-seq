package seq

import (
	"fmt"
)

func ExampleSeq_Any() {
	isEven := func(i int) bool { return i%2 == 0 }
	s1 := []int{1, 3, 5, 7, 10}
	s2 := []int{1, 3, 5, 7, 9}
	fmt.Println(FromSlice(s1).Any(isEven))
	fmt.Println(FromSlice(s2).Any(isEven))
	// Output:
	// true
	// false
}
