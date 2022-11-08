package seq

import "fmt"

func ExampleSeq_ToChan() {
	ch := make(chan int)
	go func() {
		FromSlice([]int{1, 2}).ToChan(ch)
		FromSlice([]int{3, 4}).ToChan(ch)
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}
