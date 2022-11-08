package seq

import "sync"

type zipParallelIterator[U any, V any, R any] struct {
	previous1 Iterator[U]
	previous2 Iterator[V]
	selector  func(U, V) R
	current   R
	once      sync.Once
	out       chan R
	par       int
}

type pair[U any, V any] struct {
	item1 U
	item2 V
}

func (i *zipParallelIterator[U, V, R]) run() {
	defer close(i.out)
	var wg sync.WaitGroup
	defer wg.Wait()
	in := make(chan pair[U, V])
	defer close(in)
	for j := 0; j < i.par; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				p, ok := <-in
				if !ok {
					return
				}
				i.out <- i.selector(p.item1, p.item2)
			}
		}()
	}
	for i.previous1.Next() && i.previous2.Next() {
		in <- pair[U, V]{
			item1: *i.previous1.Current(),
			item2: *i.previous2.Current(),
		}
	}
}

func (i *zipParallelIterator[U, V, R]) Next() bool {
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

func (i *zipParallelIterator[U, V, R]) Current() *R {
	return &i.current
}

// ZipP combines items of two sequences pairwise using a projection function.
// This is parallel version of Zip using specified number of goroutines to process items in parallel.
func ZipP[U any, V any, R any](par int, s1 Seq[U], s2 Seq[V], s func(U, V) R) Seq[R] {
	return seq[R]{
		iterator: func() Iterator[R] {
			return &zipParallelIterator[U, V, R]{
				previous1: s1.Iterator(),
				previous2: s2.Iterator(),
				selector:  s,
				out:       make(chan R),
				par:       par,
			}
		},
	}
}
