package seq

import (
	"fmt"
)

func ExampleDistinct() {
	s := []int{2, 2, 3, 1, 2, 3}
	r := Distinct(FromSlice(s))
	fmt.Println(r.AsSlice())
	// Output:
	// [2 3 1]
}
