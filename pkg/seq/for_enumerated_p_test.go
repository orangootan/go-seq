package seq

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSeq_ForEnumeratedP() {
	s := []string{"Huey", "Dewey", "Louie"}
	FromSlice(s).ForEnumeratedP(2, func(i int, name string) {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Println(i, name)
	})
	// Unordered output:
	// 0 Huey
	// 1 Dewey
	// 2 Louie
}

func BenchmarkSeq_ForEnumeratedP(b *testing.B) {
	par := []int{1, 2, 4}
	for _, p := range par {
		b.Run(fmt.Sprintf("par=%d", p), func(b *testing.B) {
			Repeat(1).Take(4).ForEnumeratedP(p, func(_ int, _ int) {
				time.Sleep(1 * time.Second)
			})
		})
	}
}
