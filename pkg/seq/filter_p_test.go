package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSeq_FilterP() {
	isEven := func(i int) bool {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return i%2 == 0
	}
	s := []int{1, 2, 3, 4, 5, 6}
	FromSlice(s).FilterP(2, isEven).ForEach(func(i int) {
		fmt.Println(i)
	})
	// Unordered output:
	// 2
	// 4
	// 6
}

func BenchmarkSeq_FilterP(b *testing.B) {
	isEven := func(i int) bool {
		time.Sleep(1 * time.Second)
		return i%2 == 0
	}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			FromSlice(s).FilterP(p, isEven).Count()
		})
	}
}

func ExampleSeq_WhereP() {
	isEven := func(i int) bool {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		return i%2 == 0
	}
	s := []int{1, 2, 3, 4, 5, 6}
	FromSlice(s).WhereP(2, isEven).ForEach(func(i int) {
		fmt.Println(i)
	})
	// Unordered output:
	// 2
	// 4
	// 6
}
