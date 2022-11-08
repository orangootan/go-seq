package seq

import (
	"fmt"
)

func ExampleZip() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{4, 5, 6, 7})
	r := Zip(s1, s2, func(u int, v int) [2]int {
		return [2]int{u, v}
	})
	fmt.Println(r.AsSlice())
	// Output:
	// [[1 4] [2 5] [3 6]]
}
