package seq

import "fmt"

func ExampleGroupBy() {
	isEven := func(i int) bool { return i%2 == 0 }
	s := []int{1, 2, 3, 4, 5, 6}
	g := GroupBy(FromSlice(s), isEven)
	for k, v := range g {
		fmt.Printf("%v: %v\n", k, v)
	}
	// Unordered output:
	// false: [1 3 5]
	// true: [2 4 6]
}
