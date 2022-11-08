package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleZipP() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{4, 5, 6, 7})
	ZipP(2, s1, s2, func(u int, v int) [2]int {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return [2]int{u, v}
	}).ForEach(func(s [2]int) {
		fmt.Println(s)
	})
	// Unordered output:
	// [1 4]
	// [2 5]
	// [3 6]
}

func BenchmarkZipP(b *testing.B) {
	zip := func(i int, j int) [2]int {
		time.Sleep(1 * time.Second)
		return [2]int{i, j}
	}
	s := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			ZipP(p, s, s, zip).Count()
		})
	}
}
