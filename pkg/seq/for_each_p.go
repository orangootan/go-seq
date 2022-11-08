package seq

import "sync"

// ForEachP performs an action for each item in a sequence.
// This is parallel version of ForEach using specified number of goroutines to process items in parallel.
func (s seq[T]) ForEachP(par int, do func(item T)) {
	it := s.iterator()
	in := make(chan T)
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				item, ok := <-in
				if !ok {
					return
				}
				do(item)
			}
		}()
	}
	wg.Wait()
}
