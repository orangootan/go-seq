package seq

import "fmt"

func ExampleFromChan() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	fmt.Println(FromChan(ch).AsSlice())
	// Output:
	// [1 2 3]
}
