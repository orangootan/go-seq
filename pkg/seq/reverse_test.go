package seq

import "fmt"

func ExampleSeq_Reverse() {
	s := []int{1, 2, 3, 4, 5, 6}
	r := FromSlice(s).Reverse()
	fmt.Println(r.AsSlice())
	// Output:
	// [6 5 4 3 2 1]
}
