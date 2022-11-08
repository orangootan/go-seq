package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSeq_PartitionP() {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	t, f := s.PartitionP(1, func(i int) bool {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return i%2 == 0
	})
	for _, i := range t {
		fmt.Println("even:", i)
	}
	for _, i := range f {
		fmt.Println("odd:", i)
	}
	// Unordered output:
	// even: 2
	// even: 4
	// even: 6
	// odd: 1
	// odd: 3
	// odd: 5
}

func BenchmarkSeq_PartitionP(b *testing.B) {
	isEven := func(i int) bool {
		time.Sleep(1 * time.Second)
		return i%2 == 0
	}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			FromSlice(s).PartitionP(p, isEven)
		})
	}
}
