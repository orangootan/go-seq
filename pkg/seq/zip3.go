package seq

type zip3Iterator[U any, V any, W any, R any] struct {
	previous1 Iterator[U]
	previous2 Iterator[V]
	previous3 Iterator[W]
	selector  func(U, V, W) R
}

func (i *zip3Iterator[U, V, W, R]) Next() bool {
	return i.previous1.Next() && i.previous2.Next() && i.previous3.Next()
}

func (i *zip3Iterator[U, V, W, R]) Current() *R {
	r := i.selector(*i.previous1.Current(), *i.previous2.Current(), *i.previous3.Current())
	return &r
}

// Zip3 is same as Zip but combines three sequences.
func Zip3[U any, V any, W any, R any](s1 Seq[U], s2 Seq[V], s3 Seq[W], s func(U, V, W) R) Seq[R] {
	return seq[R]{
		iterator: func() Iterator[R] {
			return &zip3Iterator[U, V, W, R]{
				previous1: s1.Iterator(),
				previous2: s2.Iterator(),
				previous3: s3.Iterator(),
				selector:  s,
			}
		},
	}
}
