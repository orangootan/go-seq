package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleMapP() {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	MapP(2, s, func(i int) int {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return i * i
	}).ForEach(func(i int) {
		fmt.Println(i)
	})
	// Unordered output:
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
}

func BenchmarkMapP(b *testing.B) {
	square := func(i int) int {
		time.Sleep(1 * time.Second)
		return i * i
	}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			MapP(p, FromSlice(s), square).Count()
		})
	}
}

func ExampleSelectP() {
	s := FromSlice([]int{1, 2, 3, 4, 5, 6})
	r := SelectP(2, s, func(i int) int {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return i * i
	})
	r.ForEach(func(i int) {
		fmt.Println(i)
	})
	// Unordered output:
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
}
