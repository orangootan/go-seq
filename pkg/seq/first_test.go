package seq

import "fmt"

func ExampleSeq_First() {
	s1 := []int{1, 2, 3}
	var s2 []int
	fmt.Println(FromSlice(s1).First())
	fmt.Println(FromSlice(s2).First())
	// Output:
	// 1 true
	// 0 false
}
