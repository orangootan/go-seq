package seq

import "fmt"

func ExampleCount() {
	s := []string{"Huey", "Dewey", "Louie"}
	c := Count[string, uint64](FromSlice(s))
	fmt.Printf("%T %v", c, c)
	// Output:
	// uint64 3
}
