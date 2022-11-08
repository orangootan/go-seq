package seq

import (
	"fmt"
)

func ExampleRepeat() {
	s := Repeat(1).Take(5)
	fmt.Println(s.AsSlice())
	// Output:
	// [1 1 1 1 1]
}
