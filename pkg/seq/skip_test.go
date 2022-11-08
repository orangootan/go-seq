package seq

import "fmt"

func ExampleSeq_Skip() {
	s := []int{1, 2, 3, 4, 5, 6}
	r := FromSlice(s).Skip(3)
	fmt.Println(r.AsSlice())
	// Output:
	// [4 5 6]
}
