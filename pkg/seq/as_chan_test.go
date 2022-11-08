package seq

import "fmt"

func ExampleSeq_AsChan() {
	ch := FromSlice([]int{1, 2, 3}).AsChan(0)
	for i := range ch {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
}
