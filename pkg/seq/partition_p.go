package seq

import "sync"

// PartitionP returns two slices. The first one contains items satisfying the given condition,
// the second one contains the rest of them.
// This is parallel version of Partition using specified number of goroutines to process items in parallel.
func (s seq[T]) PartitionP(par int, pred func(item T) bool) ([]T, []T) {
	it := s.iterator()
	in := make(chan T)
	ts := make([][]T, par)
	fs := make([][]T, par)
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				item, ok := <-in
				if !ok {
					return
				}
				if pred(item) {
					ts[j] = append(ts[j], item)
				} else {
					fs[j] = append(fs[j], item)
				}
			}
		}()
	}
	wg.Wait()
	var t, f []T
	for _, r := range ts {
		t = append(t, r...)
	}
	for _, r := range fs {
		f = append(f, r...)
	}
	return t, f
}
