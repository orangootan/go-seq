package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSeq_ForEachP() {
	s := []string{"Huey", "Dewey", "Louie"}
	FromSlice(s).ForEachP(2, func(name string) {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Println(name)
	})
	// Unordered output:
	// Huey
	// Dewey
	// Louie
}

func BenchmarkSeq_ForEachP(b *testing.B) {
	s := []int{1, 2, 3, 4}
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			FromSlice(s).ForEachP(p, func(i int) {
				time.Sleep(1 * time.Second)
			})
		})
	}
}
