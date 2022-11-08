package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSeq_AnyP() {
	isEven := func(i int) bool {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return i%2 == 0
	}
	s1 := []int{1, 3, 5, 7, 9, 11, 13, 15}
	s2 := []int{1, 2, 4, 7, 9, 11, 13, 16}
	fmt.Println(FromSlice(s1).AnyP(2, isEven))
	fmt.Println(FromSlice(s2).AnyP(2, isEven))
	// Output:
	// false
	// true
}

func BenchmarkSeq_AnyP(b *testing.B) {
	isEven := func(i int) bool {
		time.Sleep(1 * time.Second)
		return i%2 == 0
	}
	s := []int{1, 3, 5, 7, 9, 11, 13, 14}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			FromSlice(s).AnyP(p, isEven)
		})
	}
}
