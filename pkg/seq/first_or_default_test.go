package seq

import "fmt"

func ExampleSeq_FirstOrDefault() {
	s1 := []int{1, 2, 3}
	var s2 []int
	fmt.Println(FromSlice(s1).FirstOrDefault())
	fmt.Println(FromSlice(s2).FirstOrDefault())
	// Output:
	// 1
	// 0
}
