package seq

import "fmt"

func ExampleSeq_ToSlice() {
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5}
	FromSlice(s2).ToSlice(&s1)
	fmt.Println(s1)
	// Output:
	// [1 2 3 4 5]
}
