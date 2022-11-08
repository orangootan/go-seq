package seq

import (
	"fmt"
)

func ExampleGenerate() {
	s := Generate(func(i int) int { return i * i })
	fmt.Println(s.Take(6).AsSlice())
	// Output:
	// [0 1 4 9 16 25]
}
