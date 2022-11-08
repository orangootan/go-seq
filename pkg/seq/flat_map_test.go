package seq

import "fmt"

func ExampleFlatMap() {
	s := []int{1, 2, 3}
	r := FlatMap(FromSlice(s), func(i int) Seq[int] {
		return Repeat(i).Take(i)
	})
	fmt.Println(r.AsSlice())
	// Output:
	// [1 2 2 3 3 3]
}
