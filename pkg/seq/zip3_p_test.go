package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleZip3P() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{4, 5, 6})
	s3 := FromSlice([]int{7, 8, 9, 10})
	Zip3P(2, s1, s2, s3, func(u int, v int, w int) [3]int {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return [3]int{u, v, w}
	}).ForEach(func(s [3]int) {
		fmt.Println(s)
	})
	// Unordered output:
	// [1 4 7]
	// [2 5 8]
	// [3 6 9]
}

func BenchmarkZip3P(b *testing.B) {
	zip := func(i int, j int, k int) [3]int {
		time.Sleep(1 * time.Second)
		return [3]int{i, j, k}
	}
	s := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			Zip3P(p, s, s, s, zip).Count()
		})
	}
}
