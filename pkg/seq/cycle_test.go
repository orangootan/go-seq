package seq

import (
	"fmt"
)

func ExampleCycle() {
	s := []int{1, 2, 3}
	r := Cycle(s).Take(9)
	fmt.Println(r.AsSlice())
	// Output:
	// [1 2 3 1 2 3 1 2 3]
}
