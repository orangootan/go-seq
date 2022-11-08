package seq

import "sync"

type mapParallelIterator[T any, S any] struct {
	previous Iterator[T]
	selector func(item T) S
	current  S
	once     sync.Once
	out      chan S
	par      int
}

func (i *mapParallelIterator[T, S]) run() {
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
				i.out <- i.selector(item)
			}
		}()
	}
	for i.previous.Next() {
		in <- *i.previous.Current()
	}
}

func (i *mapParallelIterator[T, S]) Next() bool {
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

func (i *mapParallelIterator[T, S]) Current() *S {
	return &i.current
}

// MapP projects items of a sequence using the given function.
// This is parallel version of Map using specified number of goroutines to process items in parallel.
func MapP[T any, S any](par int, s Seq[T], fn func(item T) S) Seq[S] {
	return seq[S]{
		iterator: func() Iterator[S] {
			return &mapParallelIterator[T, S]{
				previous: s.Iterator(),
				selector: fn,
				out:      make(chan S),
				par:      par,
			}
		},
	}
}

// SelectP projects items of a sequence using the given function.
// This is parallel version of Select using specified number of goroutines to process items in parallel.
func SelectP[T any, S any](par int, s Seq[T], fn func(item T) S) Seq[S] {
	return MapP(par, s, fn)
}
