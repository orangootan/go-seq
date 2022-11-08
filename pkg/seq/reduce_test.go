package seq

import (
	"fmt"
)

func ExampleSeq_Reduce() {
	sum := func(a int, b int) int { return a + b }
	s1 := []int{1, 2, 3, 4, 5}
	var s2 []int
	fmt.Println(FromSlice(s1).Reduce(0, sum))
	fmt.Println(FromSlice(s2).Reduce(0, sum))
	// Output:
	// 15
	// 0
}
