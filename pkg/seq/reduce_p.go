package seq

import "sync"

// ReduceP applies an accumulator function over a sequence given a specified initial value.
// This is parallel version of Reduce using specified number of goroutines to process items in parallel.
func (s seq[T]) ReduceP(par int, acc T, fn func(acc T, item T) T) T {
	it := s.iterator()
	in := make(chan T)
	rs := make([]T, par)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		for it.Next() {
			in <- *it.Current()
		}
	}()
	for i := 0; i < par; i++ {
		j := i
		init, ok := <-in
		if !ok {
			break
		}
		rs[j] = init
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				item, ok := <-in
				if !ok {
					return
				}
				rs[j] = fn(rs[j], item)
			}
		}()
	}
	wg.Wait()
	for _, r := range rs {
		acc = fn(acc, r)
	}
	return acc
}
