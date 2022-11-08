package seq

import "fmt"

func ExampleSingle() {
	s := Single(3)
	fmt.Println(s.AsSlice())
	fmt.Println(s.Count())
	// Output:
	// [3]
	// 1
}
