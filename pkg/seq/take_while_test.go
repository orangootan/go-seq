package seq

import "fmt"

func ExampleSeq_TakeWhile() {
	lessThan5 := func(i int) bool { return i < 5 }
	s := []int{1, 2, 3, 4, 5, 6, 1}
	r := FromSlice(s).TakeWhile(lessThan5)
	fmt.Println(r.AsSlice())
	// Output:
	// [1 2 3 4]
}
