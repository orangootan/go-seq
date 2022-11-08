package seq

import (
	"fmt"
)

func ExampleMap() {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	r := Map(s, func(i int) int { return i * i })
	fmt.Println(r.AsSlice())
	// Output:
	// [1 4 9 16 25 36]
}

func ExampleSelect() {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	r := Select(s, func(i int) int { return i * i })
	fmt.Println(r.AsSlice())
	// Output:
	// [1 4 9 16 25 36]
}
