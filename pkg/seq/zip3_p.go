package seq

import "sync"

type zip3ParallelIterator[U any, V any, W any, R any] struct {
	previous1 Iterator[U]
	previous2 Iterator[V]
	previous3 Iterator[W]
	selector  func(U, V, W) R
	current   R
	once      sync.Once
	out       chan R
	par       int
}

type triple[U any, V any, W any] struct {
	item1 U
	item2 V
	item3 W
}

func (i *zip3ParallelIterator[U, V, W, R]) run() {
	defer close(i.out)
	var wg sync.WaitGroup
	defer wg.Wait()
	in := make(chan triple[U, V, W])
	defer close(in)
	for j := 0; j < i.par; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				t, ok := <-in
				if !ok {
					return
				}
				i.out <- i.selector(t.item1, t.item2, t.item3)
			}
		}()
	}
	for i.previous1.Next() && i.previous2.Next() && i.previous3.Next() {
		in <- triple[U, V, W]{
			item1: *i.previous1.Current(),
			item2: *i.previous2.Current(),
			item3: *i.previous3.Current(),
		}
	}
}

func (i *zip3ParallelIterator[U, V, W, R]) Next() bool {
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

func (i *zip3ParallelIterator[U, V, W, R]) Current() *R {
	return &i.current
}

// Zip3P is same as ZipP but combines three sequences.
// This is parallel version of Zip3 using specified number of goroutines to process items in parallel.
func Zip3P[U any, V any, W any, R any](par int, s1 Seq[U], s2 Seq[V], s3 Seq[W], s func(U, V, W) R) Seq[R] {
	return seq[R]{
		iterator: func() Iterator[R] {
			return &zip3ParallelIterator[U, V, W, R]{
				previous1: s1.Iterator(),
				previous2: s2.Iterator(),
				previous3: s3.Iterator(),
				selector:  s,
				out:       make(chan R),
				par:       par,
			}
		},
	}
}
