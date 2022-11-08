package seq

import "fmt"

func ExampleFromSliceReversed() {
	s := FromSliceReversed([]int{1, 2, 3, 4})
	fmt.Println(s.AsSlice())
	// Output:
	// [4 3 2 1]
}
