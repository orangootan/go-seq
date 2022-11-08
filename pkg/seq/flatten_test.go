package seq

import "fmt"

func ExampleFlatten() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{4, 5, 6})
	s3 := FromSlice([]int{7, 8, 9})
	ss := FromSlice([]Seq[int]{s1, s2, s3})
	fmt.Println(Flatten(ss).AsSlice())
	// Output:
	// [1 2 3 4 5 6 7 8 9]
}
