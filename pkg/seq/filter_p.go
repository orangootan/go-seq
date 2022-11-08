package seq

import "sync"

type filterParallelIterator[T any] struct {
	previous  Iterator[T]
	predicate func(item T) bool
	current   T
	once      sync.Once
	out       chan T
	par       int
}

func (i *filterParallelIterator[T]) run() {
	defer close(i.out)
	var wg sync.WaitGroup
	defer wg.Wait()
	in := make(chan T)
	defer close(in)
	for j := 0; j < i.par; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				item, ok := <-in
				if !ok {
					return
				}
				if i.predicate(item) {
					i.out <- item
				}
			}
		}()
	}
	for i.previous.Next() {
		in <- *i.previous.Current()
	}
}

func (i *filterParallelIterator[T]) Next() bool {
	i.once.Do(func() {
		go i.run()
	})
	item, ok := <-i.out
	if !ok {
		return false
	}
	i.current = item
	return true
}

func (i *filterParallelIterator[T]) Current() *T {
	return &i.current
}

// FilterP filters a sequence of items based on the given predicate.
// This is parallel version of Filter using specified number of goroutines to process items in parallel.
func (s seq[T]) FilterP(par int, p func(item T) bool) Seq[T] {
	return seq[T]{
		iterator: func() Iterator[T] {
			return &filterParallelIterator[T]{
				previous:  s.iterator(),
				predicate: p,
				out:       make(chan T),
				par:       par,
			}
		},
	}
}

// WhereP filters a sequence of items based on the given predicate.
// This is parallel version of Where using specified number of goroutines to process items in parallel.
func (s seq[T]) WhereP(par int, p func(item T) bool) Seq[T] {
	return s.FilterP(par, p)
}
