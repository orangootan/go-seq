package seq

import "fmt"

func ExampleSeq_SkipWhile() {
	lessThan5 := func(i int) bool { return i < 5 }
	s1 := []int{1, 2, 3, 4, 5, 6, 1}
	s2 := []int{1, 2, 3}
	r1 := FromSlice(s1).SkipWhile(lessThan5)
	r2 := FromSlice(s2).SkipWhile(lessThan5)
	fmt.Println(r1.AsSlice())
	fmt.Println(r2.AsSlice())
	// Output:
	// [5 6 1]
	// []
}
