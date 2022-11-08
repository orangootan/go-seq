package seq

import (
	"fmt"
)

func ExampleSeq_Concat() {
	s1 := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{5, 6, 7})
	r := s1.Concat(s2)
	fmt.Println(r.AsSlice())
	// Output:
	// [1 2 3 4 5 6 7]
}
