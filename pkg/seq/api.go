package seq

// Seq is the main interface declaring methods for working with sequences.
type Seq[T any] interface {
	// Iterator returns iterator object of a sequence.
	Iterator() Iterator[T]

	// All determines whether all items of a sequence satisfy the given condition.
	All(p func(item T) bool) bool

	// Any determines whether any item of a sequence satisfies the given condition.
	Any(p func(item T) bool) bool

	// AsChan creates a channel and sends all sequence items into it in separate goroutine.
	// After the sequence exhausted the channel is closed.
	// You can specify the channel capacity with cap parameter (can be 0).
	AsChan(cap int) <-chan T

	// AsSlice collects all sequence items into newly created slice.
	AsSlice() []T

	// Concat concatenates two sequences.
	Concat(s2 Seq[T]) Seq[T]

	// Count counts number of items in a sequence.
	Count() int

	// Cycled repeats a sequence infinitely.
	Cycled() Seq[T]

	// Filter filters a sequence of items based on the given predicate.
	Filter(p func(item T) bool) Seq[T]

	// First returns undefined value and false if a sequence is empty
	// or the first item and true otherwise.
	First() (T, bool)

	// FirstOrDefault returns default value of type T if a sequence is empty
	// or the first item otherwise.
	FirstOrDefault() T

	// ForEach executes the given function for each item in a sequence.
	ForEach(do func(item T))

	// ForEnumerated is similar to ForEach but produces ordinal numbers of items starting from 0
	// as the first parameter of the given function.
	ForEnumerated(do func(i int, item T))

	// Inspect allows to observe items of a sequence by executing a function for each item.
	Inspect(do func(item T)) Seq[T]

	// Interleave produces a sequence that interleaves items from two sequences.
	Interleave(s2 Seq[T]) Seq[T]

	// IsEmpty determines whether a sequence is empty.
	IsEmpty() bool

	// Partition returns two slices. The first one contains items satisfying the given condition,
	// the second one contains the rest of them.
	Partition(p func(item T) bool) ([]T, []T)

	// Reduce applies the given accumulator function over a sequence given a specified initial value.
	Reduce(acc T, fn func(acc T, item T) T) T

	// Reverse inverts the order of items in a sequence.
	Reverse() Seq[T]

	// Skip bypasses the given number of items in a sequence and then returns the remaining items.
	Skip(n int) Seq[T]

	// SkipWhile bypasses items in a sequence as long as the given condition is true
	// and then returns the remaining items.
	SkipWhile(p func(item T) bool) Seq[T]

	// Take returns the given number of items from the start of a sequence.
	Take(n int) Seq[T]

	// TakeWhile returns items from a sequence as long as the given condition is true,
	// and then skips the remaining items.
	TakeWhile(p func(item T) bool) Seq[T]

	// ToChan sends sequence items to an existing channel in current goroutine so this call is blocking.
	// This function does not close the channel after the sequence is exhausted.
	ToChan(ch chan<- T)

	// ToSlice appends items of a sequence to an existing slice given by reference.
	ToSlice(acc *[]T)

	// Where filters a sequence of items based on the given predicate.
	Where(p func(item T) bool) Seq[T]

	// AllP determines whether all items of a sequence satisfy the given condition.
	// This is parallel version of All using specified number of goroutines to process items in parallel.
	AllP(par int, p func(item T) bool) bool

	// AnyP determines whether any item of a sequence satisfies the given condition.
	// This is parallel version of Any using specified number of goroutines to process items in parallel.
	AnyP(par int, p func(item T) bool) bool

	// FilterP filters a sequence of items based on the given predicate.
	// This is parallel version of Filter using specified number of goroutines to process items in parallel.
	FilterP(par int, p func(item T) bool) Seq[T]

	// ForEachP performs an action for each item in a sequence.
	// This is parallel version of ForEach using specified number of goroutines to process items in parallel.
	ForEachP(par int, do func(item T))

	// ForEnumeratedP is similar to ForEachP but passes item index as first parameter of action function.
	// This is parallel version of ForEnumerated using specified number of goroutines to process items in parallel.
	ForEnumeratedP(par int, do func(i int, item T))

	// PartitionP returns two slices. The first one contains items satisfying the given condition,
	// the second one contains the rest of them.
	// This is parallel version of Partition using specified number of goroutines to process items in parallel.
	PartitionP(par int, p func(item T) bool) ([]T, []T)

	// ReduceP applies an accumulator function over a sequence given a specified initial value.
	// This is parallel version of Reduce using specified number of goroutines to process items in parallel.
	ReduceP(par int, acc T, fn func(acc T, item T) T) T

	// WhereP filters a sequence of items based on the given predicate.
	// This is parallel version of Where using specified number of goroutines to process items in parallel.
	WhereP(par int, p func(item T) bool) Seq[T]
}
