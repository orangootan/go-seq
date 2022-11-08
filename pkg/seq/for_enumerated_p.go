package seq

import "sync"

type enumerated[T any] struct {
	index int
	item  T
}

// ForEnumeratedP is similar to ForEachP but passes item index as first parameter of action function.
// This is parallel version of ForEnumerated using specified number of goroutines to process items in parallel.
func (s seq[T]) ForEnumeratedP(par int, do func(i int, item T)) {
	it := s.iterator()
	in := make(chan enumerated[T])
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		i := -1
		for it.Next() {
			i++
			in <- enumerated[T]{
				index: i,
				item:  *it.Current(),
			}
		}
	}()
	for i := 0; i < par; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				e, ok := <-in
				if !ok {
					return
				}
				do(e.index, e.item)
			}
		}()
	}
	wg.Wait()
}
