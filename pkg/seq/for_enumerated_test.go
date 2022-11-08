package seq

import "fmt"

func ExampleSeq_ForEnumerated() {
	s := []string{"Huey", "Dewey", "Louie"}
	FromSlice(s).ForEnumerated(func(i int, name string) {
		fmt.Println(i, name)
	})
	// Output:
	// 0 Huey
	// 1 Dewey
	// 2 Louie
}
