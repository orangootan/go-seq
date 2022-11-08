package seq

import "fmt"

func ExampleSeq_Cycled() {
	s := []int{1, 2, 3}
	r := FromSlice(s).Cycled().Take(10)
	fmt.Println(r.AsSlice())
	// Output:
	// [1 2 3 1 2 3 1 2 3 1]
}
