package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleFlatMapP() {
	s := FromSlice([]int{1, 2, 3})
	FlatMapP(2, s, func(i int) Seq[int] {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return Repeat(i).Take(i)
	}).ForEach(func(i int) {
		fmt.Println(i)
	})
	// Unordered output:
	// 1
	// 2
	// 2
	// 3
	// 3
	// 3
}

func BenchmarkFlatMapP(b *testing.B) {
	fn := func(i int) Seq[int] {
		time.Sleep(1 * time.Second)
		return Repeat(i).Take(i)
	}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			FlatMapP(p, FromSlice(s), fn).Count()
		})
	}
}
