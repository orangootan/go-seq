package seq

import (
	"fmt"
)

func ExampleEqual() {
	s1 := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{1, 2, 3})
	s3 := FromSlice([]int{0, 1, 2})
	s4 := FromSlice([]int{0, 1, 2})
	fmt.Println(Equal(s1, s2))
	fmt.Println(Equal(s2, s3))
	fmt.Println(Equal(s3, s4))
	// Output:
	// false
	// false
	// true
}
