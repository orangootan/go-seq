package seq

import "fmt"

func ExampleSeq_Iterator() {
	s := []int{1, 2, 3}
	it := FromSlice(s).Iterator()
	for it.Next() {
		fmt.Println(*it.Current())
	}
	// Output:
	// 1
	// 2
	// 3
}
