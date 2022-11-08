package seq

import (
	"fmt"
)

func ExampleSeq_Filter() {
	isEven := func(i int) bool { return i%2 == 0 }
	s := []int{1, 2, 3, 4, 5, 6}
	r := FromSlice(s).Filter(isEven)
	fmt.Println(r.AsSlice())
	// Output:
	// [2 4 6]
}

func ExampleSeq_Where() {
	isEven := func(i int) bool { return i%2 == 0 }
	s := []int{1, 2, 3, 4, 5, 6}
	r := FromSlice(s).Where(isEven)
	fmt.Println(r.AsSlice())
	// Output:
	// [2 4 6]
}
