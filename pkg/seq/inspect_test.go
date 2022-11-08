package seq

import "fmt"

func ExampleSeq_Inspect() {
	isEven := func(i int) bool { return i%2 == 0 }
	FromSlice([]int{1, 2, 3}).
		Inspect(func(i int) {
			fmt.Printf("before: %v\n", i)
		}).
		Filter(isEven).
		Inspect(func(i int) {
			fmt.Printf("after: %v\n", i)
		}).
		Count()
	// Output:
	// before: 1
	// before: 2
	// after: 2
	// before: 3
}
