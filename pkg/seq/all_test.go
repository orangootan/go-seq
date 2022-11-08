package seq

import "fmt"

func ExampleSeq_All() {
	isEven := func(i int) bool { return i%2 == 0 }
	s1 := []int{2, 4, 6, 8, 10}
	s2 := []int{2, 4, 6, 8, 11}
	fmt.Println(FromSlice(s1).All(isEven))
	fmt.Println(FromSlice(s2).All(isEven))
	// Output:
	// true
	// false
}
