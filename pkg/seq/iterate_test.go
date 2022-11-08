package seq

import "fmt"

func ExampleIterate() {
	s := Iterate(2, func(i int) int { return 2 * i })
	fmt.Println(s.Take(8).AsSlice())
	// Output:
	// [2 4 8 16 32 64 128 256]
}
