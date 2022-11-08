package seq

import (
	"fmt"
)

func ExampleFold() {
	sum := func(a int, b int) int { return a + b }
	appendInt := func(s []int, i int) []int {
		return append(s, i)
	}
	s := FromSlice([]int{1, 2, 3})
	fmt.Println(Fold(s, 0, sum))
	fmt.Println(Fold(s, []int{4, 5}, appendInt))
	// Output:
	// 6
	// [4 5 1 2 3]
}
