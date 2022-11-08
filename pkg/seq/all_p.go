package seq

import (
	"sync"
)

// AllP determines whether all items of a sequence satisfy the given condition.
// This is parallel version of All using specified number of goroutines to process items in parallel.
func (s seq[T]) AllP(par int, p func(item T) bool) bool {
	it := s.iterator()
	in := make(chan T)
	stop := make(chan int)
	var wg sync.WaitGroup
	var once sync.Once
	result := true
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		for it.Next() {
			select {
			case in <- *it.Current():
			case _, _ = <-stop:
				return
			}
		}
	}()
	for i := 0; i < par; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				item, ok := <-in
				if !ok {
					return
				}
				if !p(item) {
					once.Do(func() {
						close(stop)
						result = false
					})
				}
			}
		}()
	}
	wg.Wait()
	return result
}
