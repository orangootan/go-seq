package seq

import "fmt"

func ExampleSeq_IsEmpty() {
	s1 := []int{1, 2, 3}
	var s2 []int
	fmt.Println(FromSlice(s1).IsEmpty())
	fmt.Println(FromSlice(s2).IsEmpty())
	fmt.Println(Empty[int]().IsEmpty())
	// Output:
	// false
	// true
	// true
}
