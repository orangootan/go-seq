package seq

import "fmt"

func ExampleSeq_Partition() {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	t, f := s.Partition(func(i int) bool { return i%2 == 0 })
	fmt.Println(t)
	fmt.Println(f)
	// Output:
	// [2 4 6]
	// [1 3 5]
}
