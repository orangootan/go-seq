package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSeq_AllP() {
	isEven := func(i int) bool {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return i%2 == 0
	}
	s1 := []int{2, 4, 6, 8, 10, 12, 14, 16}
	s2 := []int{2, 4, 6, 8, 11, 12, 14, 16}
	fmt.Println(FromSlice(s1).AllP(2, isEven))
	fmt.Println(FromSlice(s2).AllP(2, isEven))
	// Output:
	// true
	// false
}

func BenchmarkSeq_AllP(b *testing.B) {
	isEven := func(i int) bool {
		time.Sleep(1 * time.Second)
		return i%2 == 0
	}
	s := []int{2, 4, 6, 8, 10, 12, 14, 17}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			FromSlice(s).AllP(p, isEven)
		})
	}
}
