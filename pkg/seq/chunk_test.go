package seq

import "fmt"

func ExampleChunk() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	r := Chunk(3, FromSlice(s))
	fmt.Println(r.AsSlice())
	// Output:
	// [[1 2 3] [4 5 6] [7 8 9] [10 11]]
}
