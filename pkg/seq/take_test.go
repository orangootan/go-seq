package seq

import "fmt"

func ExampleSeq_Take() {
	s := []int{1, 2, 3, 4, 5, 6}
	r := FromSlice(s).Take(4)
	fmt.Println(r.AsSlice())
	// Output:
	// [1 2 3 4]
}
