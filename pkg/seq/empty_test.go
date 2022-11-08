package seq

import "fmt"

func ExampleEmpty() {
	fmt.Println(Empty[int]().Count())
	fmt.Println(Empty[string]().AsSlice())
	fmt.Println(Empty[int]().IsEmpty())
	it := Empty[int]().Iterator()
	fmt.Println(it.Next(), it.Current())
	// Output:
	// 0
	// []
	// true
	// false <nil>
}
