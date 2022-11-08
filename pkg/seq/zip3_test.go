package seq

import (
	"fmt"
)

func ExampleZip3() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{4, 5, 6})
	s3 := FromSlice([]int{7, 8, 9, 10})
	r := Zip3(s1, s2, s3, func(u int, v int, w int) [3]int {
		return [3]int{u, v, w}
	})
	fmt.Println(r.AsSlice())
	// Output:
	// [[1 4 7] [2 5 8] [3 6 9]]
}
