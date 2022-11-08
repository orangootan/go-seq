package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSeq_ReduceP() {
	sum := func(a int, b int) int {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return a + b
	}
	s1 := []int{1, 2, 3, 4, 5}
	var s2 []int
	fmt.Println(FromSlice(s1).ReduceP(2, 0, sum))
	fmt.Println(FromSlice(s2).ReduceP(2, 0, sum))
	// Output:
	// 15
	// 0
}

func BenchmarkSeq_ReduceP(b *testing.B) {
	inc := func(i int) int { return i + 1 }
	sum := func(a int, b int) int {
		time.Sleep(1 * time.Second)
		return a + b
	}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			Iterate(1, inc).Take(20).ReduceP(p, 0, sum)
		})
	}
}
