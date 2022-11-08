package seq

import "fmt"

func ExampleSeq_ForEach() {
	s := []string{"Huey", "Dewey", "Louie"}
	FromSlice(s).ForEach(func(name string) {
		fmt.Println(name)
	})
	// Output:
	// Huey
	// Dewey
	// Louie
}
