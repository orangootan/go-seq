package seq

import (
	"fmt"
)

func ExampleSeq_Count() {
	s1 := []int{1, 2, 3, 4, 5}
	var s2 []int
	fmt.Println(FromSlice(s1).Count())
	fmt.Println(FromSlice(s2).Count())
	// Output:
	// 5
	// 0
}
