package seq

import "fmt"

func ExampleSeq_Interleave() {
	s1 := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{4, 5, 6})
	r := s1.Interleave(s2)
	fmt.Println(r.AsSlice())
	// Output:
	// [1 4 2 5 3 6 4]
}
