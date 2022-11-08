package seq

type zipIterator[U any, V any, R any] struct {
	previous1 Iterator[U]
	previous2 Iterator[V]
	selector  func(U, V) R
}

func (i *zipIterator[U, V, R]) Next() bool {
	return i.previous1.Next() && i.previous2.Next()
}

func (i *zipIterator[U, V, R]) Current() *R {
	r := i.selector(*i.previous1.Current(), *i.previous2.Current())
	return &r
}

// Zip combines items of two sequences pairwise using a projection function.
func Zip[U any, V any, R any](s1 Seq[U], s2 Seq[V], s func(U, V) R) Seq[R] {
	return seq[R]{
		iterator: func() Iterator[R] {
			return &zipIterator[U, V, R]{
				previous1: s1.Iterator(),
				previous2: s2.Iterator(),
				selector:  s,
			}
		},
	}
}
