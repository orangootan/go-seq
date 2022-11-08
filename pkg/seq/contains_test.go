package seq

import (
	"fmt"
)

func ExampleContains() {
	s := FromSlice([]int{1, 2, 3, 4, 5})
	fmt.Println(Contains(3, s))
	fmt.Println(Contains(6, s))
	// Output:
	// true
	// false
}
