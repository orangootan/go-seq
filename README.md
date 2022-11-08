[![License: MIT](https://img.shields.io/badge/License-MIT-pink.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/orangootan/go-seq)]()
[![Release](https://img.shields.io/github/release/orangootan/go-seq)](https://github.com/orangootan/go-seq/releases/latest)
[![GoDoc](https://godoc.org/github.com/orangootan/go-seq?status.svg)](https://godoc.org/github.com/orangootan/go-seq)
[![Go Report Card](https://goreportcard.com/badge/github.com/orangootan/go-seq)](https://goreportcard.com/report/github.com/orangootan/go-seq)
[![CodeFactor](https://www.codefactor.io/repository/github/orangootan/go-seq/badge)](https://www.codefactor.io/repository/github/orangootan/go-seq)
![Test](https://github.com/orangootan/go-seq/actions/workflows/test.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/orangootan/go-seq/badge.svg?branch=main)](https://coveralls.io/github/orangootan/go-seq?branch=main)
[![codecov](https://codecov.io/gh/orangootan/go-seq/branch/main/graph/badge.svg?token=EFWGDYVM0Y)](https://codecov.io/gh/orangootan/go-seq)
## go-seq
Golang implementation of generic lazily iterated sequences.  
Requires Go 1.18+.
## Package
```go
import "github.com/orangootan/go-seq/pkg/seq"
```
All examples use unqualified import as if:
```go
import . "github.com/orangootan/go-seq/pkg/seq"
```
## Examples
### Seq[T]
**Seq[T]** is the main interface declaring methods for working with sequences. Some functions are implemented as standalone functions (not methods of this interface) due to limitations of Go generics implementation.
### FromSlice
**FromSlice()** returns a sequence of items from the given slice.  
See further examples.
### Iterator
**Iterator()** returns iterator object of a sequence.  
**Next()** advances iterator to the next item of the iterated sequence. Returns true if successfully advanced or false if passed the end of the sequence.  
**Current()** returns pointer to the current item of the iterated sequence.
```go
s := []int{1, 2, 3}
it := FromSlice(s).Iterator()
for it.Next() {
    fmt.Println(*it.Current())
}
// Output:
// 1
// 2
// 3
```
### AsSlice
**AsSlice()** collects all sequence items into newly created slice.  
See further examples.
### FromSliceReversed
**FromSliceReversed()** returns a sequence of items from the given slice in inverted order.
```go
s := FromSliceReversed([]int{1, 2, 3, 4})
fmt.Println(s.AsSlice())
// Output:
// [4 3 2 1]
```
### FromChan
**FromChan()** returns a sequence of items received from a channel until the channel is closed.
```go
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()
fmt.Println(FromChan(ch).AsSlice())
// Output:
// [1 2 3]
```
### Generate
**Generate()** returns an infinite sequence applying a mapping function to item indices.
```go
s := Generate(func(i int) int { return i * i })
fmt.Println(s.Take(6).AsSlice())
// Output:
// [0 1 4 9 16 25]
```
### Iterate
**Iterate()** returns an infinite sequence consisting of seed, fn(seed), fn(fn(seed)), etc.
```go
s := Iterate(2, func(i int) int { return 2 * i })
fmt.Println(s.Take(8).AsSlice())
// Output:
// [2 4 8 16 32 64 128 256]
```
### Cycle
**Cycle()** returns an infinite sequence by cycling items from the given slice.
```go
s := []int{1, 2, 3}
r := Cycle(s).Take(9)
fmt.Println(r.AsSlice())
// Output:
// [1 2 3 1 2 3 1 2 3]
```
### Repeat
**Repeat()** returns an infinite sequence repeating the given item.
```go
s := Repeat(1).Take(5)
fmt.Println(s.AsSlice())
// Output:
// [1 1 1 1 1]
```
### Single
// Single returns a sequence containing exactly one specified item.
```go
s := Single(3)
fmt.Println(s.AsSlice())
fmt.Println(s.Count())
// Output:
// [3]
// 1
```
### Empty
**Empty()** returns empty sequence.
```go
fmt.Println(Empty[int]().Count())
fmt.Println(Empty[string]().AsSlice())
fmt.Println(Empty[int]().IsEmpty())
it := Empty[int]().Iterator()
fmt.Println(it.Next(), it.Current())
// Output:
// 0
// []
// true
// false <nil>
```
### IsEmpty
**IsEmpty()** determines whether a sequence is empty.
```go
s1 := []int{1, 2, 3}
var s2 []int
fmt.Println(FromSlice(s1).IsEmpty())
fmt.Println(FromSlice(s2).IsEmpty())
fmt.Println(Empty[int]().IsEmpty())
// Output:
// false
// true
// true
```
### Equal
**Equal()** determines whether two sequences have equal length and equal corresponding items. This function works only with sequences of comparable items.
```go
s1 := FromSlice([]int{1, 2, 3, 4})
s2 := FromSlice([]int{1, 2, 3})
s3 := FromSlice([]int{0, 1, 2})
s4 := FromSlice([]int{0, 1, 2})
fmt.Println(Equal(s1, s2))
fmt.Println(Equal(s2, s3))
fmt.Println(Equal(s3, s4))
// Output:
// false
// false
// true
```
### Concat
**Concat()** concatenates two sequences.
```go
s1 := FromSlice([]int{1, 2, 3, 4})
s2 := FromSlice([]int{5, 6, 7})
r := s1.Concat(s2)
fmt.Println(r.AsSlice())
// Output:
// [1 2 3 4 5 6 7]
```
### Cycled
**Cycled()** repeats a sequence infinitely.
```go
s := []int{1, 2, 3}
r := FromSlice(s).Cycled().Take(10)
fmt.Println(r.AsSlice())
// Output:
// [1 2 3 1 2 3 1 2 3 1]
```
### Filter (Where)
**Filter()** filters a sequence of items based on the given predicate.  
**Where()** is an alias for **Filter()**. Use whatever name you like.
```go
isEven := func(i int) bool { return i%2 == 0 }
s := []int{1, 2, 3, 4, 5, 6}
r := FromSlice(s).Filter(isEven)
fmt.Println(r.AsSlice())
// Output:
// [2 4 6]
```
### Reverse
**Reverse()** inverts the order of items in a sequence.
```go
s := []int{1, 2, 3, 4, 5, 6}
r := FromSlice(s).Reverse()
fmt.Println(r.AsSlice())
// Output:
// [6 5 4 3 2 1]
```
### Skip
**Skip()** bypasses the given number of items in a sequence and then returns the remaining items.
```go
s := []int{1, 2, 3, 4, 5, 6}
r := FromSlice(s).Skip(3)
fmt.Println(r.AsSlice())
// Output:
// [4 5 6]
```
### SkipWhile
**SkipWhile()** bypasses items in a sequence as long as the given condition is true and then returns the remaining items.
```go
lessThan5 := func(i int) bool { return i < 5 }
s1 := []int{1, 2, 3, 4, 5, 6, 1}
s2 := []int{1, 2, 3}
r1 := FromSlice(s1).SkipWhile(lessThan5)
r2 := FromSlice(s2).SkipWhile(lessThan5)
fmt.Println(r1.AsSlice())
fmt.Println(r2.AsSlice())
// Output:
// [5 6 1]
// []
```
### Take
**Take()** returns the given number of items from the start of a sequence.
```go
s := []int{1, 2, 3, 4, 5, 6}
r := FromSlice(s).Take(4)
fmt.Println(r.AsSlice())
// Output:
// [1 2 3 4]
```
### TakeWhile
**TakeWhile()** returns items from a sequence as long as the given condition is true, and then skips the remaining items.
```go
lessThan5 := func(i int) bool { return i < 5 }
s := []int{1, 2, 3, 4, 5, 6, 1}
r := FromSlice(s).TakeWhile(lessThan5)
fmt.Println(r.AsSlice())
// Output:
// [1 2 3 4]
```
### Interleave
**Interleave()** returns a sequence that interleaves items from two sequences.
```go
s1 := FromSlice([]int{1, 2, 3, 4})
s2 := FromSlice([]int{4, 5, 6})
r := s1.Interleave(s2)
fmt.Println(r.AsSlice())
// Output:
// [1 4 2 5 3 6 4]
```
### Map (Select)
**Map()** projects items of a sequence using the given function.  
**Select()** is an alias for **Map()**. Use whatever name you like.
```go
s := FromSlice([]int{1, 2, 3, 4, 5, 6})
r := Map(s, func(i int) int { return i * i })
fmt.Println(r.AsSlice())
// Output:
// [1 4 9 16 25 36]
```
### Zip
**Zip()** combines items of two sequences pairwise using a projection function.
```go
s1 := FromSlice([]int{1, 2, 3})
s2 := FromSlice([]int{4, 5, 6, 7})
r := Zip(s1, s2, func(u int, v int) [2]int {
    return [2]int{u, v}
})
fmt.Println(r.AsSlice())
// Output:
// [[1 4] [2 5] [3 6]]
```
### Zip3
**Zip3()** is same as Zip but combines three sequences.
```go
s1 := FromSlice([]int{1, 2, 3})
s2 := FromSlice([]int{4, 5, 6})
s3 := FromSlice([]int{7, 8, 9, 10})
r := Zip3(s1, s2, s3, func(u int, v int, w int) [3]int {
    return [3]int{u, v, w}
})
fmt.Println(r.AsSlice())
// Output:
// [[1 4 7] [2 5 8] [3 6 9]]
```
### All
**All()** determines whether all items of a sequence satisfy the given condition.
```go
isEven := func(i int) bool { return i%2 == 0 }
s1 := []int{2, 4, 6, 8, 10}
s2 := []int{2, 4, 6, 8, 11}
fmt.Println(FromSlice(s1).All(isEven))
fmt.Println(FromSlice(s2).All(isEven))
// Output:
// true
// false
```
### Any
**Any()** determines whether any item of a sequence satisfies the given condition.
```go
isEven := func(i int) bool { return i%2 == 0 }
s1 := []int{1, 3, 5, 7, 10}
s2 := []int{1, 3, 5, 7, 9}
fmt.Println(FromSlice(s1).Any(isEven))
fmt.Println(FromSlice(s2).Any(isEven))
// Output:
// true
// false
```
### Chunk
**Chunk()** splits a sequence into chunks of the given size and returns sequence of slices.
```go
s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
r := Chunk(3, FromSlice(s))
fmt.Println(r.AsSlice())
// Output:
// [[1 2 3] [4 5 6] [7 8 9] [10 11]]
```
### Contains
**Contains()** determines whether a sequence contains the given item. This function works only with sequences of comparable items.
```go
s := FromSlice([]int{1, 2, 3, 4, 5})
fmt.Println(Contains(3, s))
fmt.Println(Contains(6, s))
// Output:
// true
// false
```
### Count (method)
**Count()** counts number of items in a sequence.
```go
s1 := []int{1, 2, 3, 4, 5}
var s2 []int
fmt.Println(FromSlice(s1).Count())
fmt.Println(FromSlice(s2).Count())
// Output:
// 5
// 0
```
### Count (function)
**Count()** counts number of items in a sequence. You can specify an integer type for result.
```go
s := []string{"Huey", "Dewey", "Louie"}
c := Count[string, uint64](FromSlice(s))
fmt.Printf("%T %v", c, c)
// Output:
// uint64 3
```
### Distinct
**Distinct()** returns distinct elements from a sequence. This function works only with sequences of comparable items.
```go
s := []int{2, 2, 3, 1, 2, 3}
r := Distinct(FromSlice(s))
fmt.Println(r.AsSlice())
// Output:
// [2 3 1]
```
### First
**First()** returns undefined value and false if a sequence is empty or the first item and true otherwise.
```go
s1 := []int{1, 2, 3}
var s2 []int
fmt.Println(FromSlice(s1).First())
fmt.Println(FromSlice(s2).First())
// Output:
// 1 true
// 0 false
```
### FirstOrDefault
**FirstOrDefault()** returns default value of type T if a sequence is empty or the first item otherwise.
```go
s1 := []int{1, 2, 3}
var s2 []int
fmt.Println(FromSlice(s1).FirstOrDefault())
fmt.Println(FromSlice(s2).FirstOrDefault())
// Output:
// 1
// 0
```
### Flatten
**Flatten()** combines a sequence of item sequences into flat sequence of items.
```go
s1 := FromSlice([]int{1, 2, 3})
s2 := FromSlice([]int{4, 5, 6})
s3 := FromSlice([]int{7, 8, 9})
ss := FromSlice([]Seq[int]{s1, s2, s3})
fmt.Println(Flatten(ss).AsSlice())
// Output:
// [1 2 3 4 5 6 7 8 9]
```
### FlatMap
**FlatMap()** maps each sequence item to a sequence and flattens results.
```go
s := []int{1, 2, 3}
r := FlatMap(FromSlice(s), func(i int) Seq[int] {
    return Repeat(i).Take(i)
})
fmt.Println(r.AsSlice())
// Output:
// [1 2 2 3 3 3]
```
### Reduce
**Reduce()** applies an accumulator function over a sequence given a specified initial value.
```go
sum := func(a int, b int) int { return a + b }
s1 := []int{1, 2, 3, 4, 5}
var s2 []int
fmt.Println(FromSlice(s1).Reduce(0, sum))
fmt.Println(FromSlice(s2).Reduce(0, sum))
// Output:
// 15
// 0
```
### Fold
**Fold()** applies an accumulator function over a sequence given a specified initial value. Accumulator and item types can be different.
```go
sum := func(a int, b int) int { return a + b }
appendInt := func(s []int, i int) []int {
    return append(s, i)
}
s := FromSlice([]int{1, 2, 3})
fmt.Println(Fold(s, 0, sum))
fmt.Println(Fold(s, []int{4, 5}, appendInt))
// Output:
// 6
// [4 5 1 2 3]
```
### ForEach
**ForEach()** performs an action for each item in a sequence.
```go
s := []string{"Huey", "Dewey", "Louie"}
FromSlice(s).ForEach(func(name string) {
    fmt.Println(name)
})
// Output:
// Huey
// Dewey
// Louie
```
### ForEnumerated
**ForEnumerated()** is similar to ForEach but passes item index as first parameter of action function.
```go
s := []string{"Huey", "Dewey", "Louie"}
FromSlice(s).ForEnumerated(func(i int, name string) {
    fmt.Println(i, name)
})
// Output:
// 0 Huey
// 1 Dewey
// 2 Louie
```
### GroupBy
**GroupBy()** returns a map grouping sequence items by key produced by the given function. Key type must be comparable.
```go
isEven := func(i int) bool { return i%2 == 0 }
s := []int{1, 2, 3, 4, 5, 6}
g := GroupBy(FromSlice(s), isEven)
for k, v := range g {
    fmt.Printf("%v: %v\n", k, v)
}
// Unordered output:
// false: [1 3 5]
// true: [2 4 6]
```
### Inspect
**Inspect()** allows to observe items of a sequence by performing an action for each item.
```go
isEven := func(i int) bool { return i%2 == 0 }
FromSlice([]int{1, 2, 3}).
    Inspect(func(i int) {
        fmt.Printf("before: %v\n", i)
    }).
    Filter(isEven).
    Inspect(func(i int) {
        fmt.Printf("after: %v\n", i)
    }).
    Count()
// Output:
// before: 1
// before: 2
// after: 2
// before: 3
```
### Partition
**Partition()** returns two slices. The first one contains items satisfying the given condition, the second one contains the rest of them.
```go
s := FromSlice([]int{1, 2, 3, 4, 5, 6})
t, f := s.Partition(func(i int) bool { return i%2 == 0 })
fmt.Println(t)
fmt.Println(f)
// Output:
// [2 4 6]
// [1 3 5]
```
### Sum
**Sum()** returns sum of elements in a sequence. This function works with sequences of numbers.
```go
s1 := []int{1, 2, 3, 4}
var s2 []int
fmt.Println(Sum(FromSlice(s1)))
fmt.Println(Sum(FromSlice(s2)))
// Output:
// 10
// 0
```
### Product
**Product()** returns product of elements in a sequence. This function works with sequences of numbers.
```go
s1 := []int{1, 2, 3, 4}
var s2 []int
fmt.Println(Product(FromSlice(s1)))
fmt.Println(Product(FromSlice(s2)))
// Output:
// 24
// 1
```
### Average
**Average()** returns average number of a sequence of numbers. This function works with sequences of integer and float numbers.
```go
s1 := []int{1, 2, 3, 4, 5}
fmt.Println(Average(FromSlice(s1)))
// Output:
// 3
```
### Max
**Max()** returns the maximum element of a sequence and true if the sequence is not empty or undefined value and false otherwise.  This function works with sequences of integer numbers, float numbers and strings.
```go
s1 := []int{1, 2, 3}
s2 := []string{"Huey", "Dewey", "Louie"}
var s3 []int
fmt.Println(Max(FromSlice(s1)))
fmt.Println(Max(FromSlice(s2)))
fmt.Println(Max(FromSlice(s3)))
// Output:
// 3 true
// Louie true
// 0 false
```
### Min
**Min()** returns the minimum element of a sequence and true if the sequence is not empty or undefined value and false otherwise. This function works with sequences of integer numbers, float numbers and strings.
```go
s1 := []int{1, 2, 3}
s2 := []string{"Huey", "Dewey", "Louie"}
var s3 []int
fmt.Println(Min(FromSlice(s1)))
fmt.Println(Min(FromSlice(s2)))
fmt.Println(Min(FromSlice(s3)))
// Output:
// 1 true
// Dewey true
// 0 false
```
### AsChan
**AsChan()** creates a channel and sends all sequence items into it in separate goroutine. After the sequence exhausted the channel is closed. You can specify the channel capacity with cap parameter.
```go
ch := FromSlice([]int{1, 2, 3}).AsChan(0)
for i := range ch {
    fmt.Println(i)
}
// Output:
// 1
// 2
// 3
```
### ToChan
**ToChan()** sends sequence items to an existing channel in current goroutine so this call is blocking. This function does not close the channel after the sequence is exhausted.
```go
ch := make(chan int)
go func() {
    FromSlice([]int{1, 2}).ToChan(ch)
    FromSlice([]int{3, 4}).ToChan(ch)
    close(ch)
}()
for i := range ch {
    fmt.Println(i)
}
// Output:
// 1
// 2
// 3
// 4
```
### ToSlice
**ToSlice()** appends items of a sequence to an existing slice given by reference.
```go
s1 := []int{1, 2}
s2 := []int{3, 4, 5}
FromSlice(s2).ToSlice(&s1)
fmt.Println(s1)
// Output:
// [1 2 3 4 5]
```
### FilterP (WhereP)
**FilterP()** filters a sequence of items based on the given predicate.  
This is parallel version of **Filter()** using specified number of goroutines to process items in parallel.
**WhereP()** is an alias for **FilterP()**. Use whatever name you like.
```go
isEven := func(i int) bool {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return i%2 == 0
}
s := []int{1, 2, 3, 4, 5, 6}
FromSlice(s).FilterP(2, isEven).ForEach(func(i int) {
    fmt.Println(i)
})
// Unordered output:
// 2
// 4
// 6
```
### MapP (SelectP)
**MapP()** projects items of a sequence using the given function.  
This is parallel version of **Map()** using specified number of goroutines to process items in parallel.
**SelectP()** is an alias for **MapP()**. Use whatever name you like.
```go
s := FromSlice([]int{1, 2, 3, 4, 5, 6})
MapP(2, s, func(i int) int {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return i * i
}).ForEach(func(i int) {
    fmt.Println(i)
})
// Unordered output:
// 1
// 4
// 9
// 16
// 25
// 36
```
### FlatmapP
**FlatMapP()** maps each sequence item to a sequence and flattens results.  
This is parallel version of **FlatMap()** using specified number of goroutines to process items in parallel.
```go
s := FromSlice([]int{1, 2, 3})
FlatMapP(2, s, func(i int) Seq[int] {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return Repeat(i).Take(i)
}).ForEach(func(i int) {
    fmt.Println(i)
})
// Unordered output:
// 1
// 2
// 2
// 3
// 3
// 3
```
### ReduceP
**ReduceP()** applies an accumulator function over a sequence given a specified initial value.  
This is parallel version of **Reduce()** using specified number of goroutines to process items in parallel.
```go
sum := func(a int, b int) int {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return a + b
}
s1 := []int{1, 2, 3, 4, 5}
var s2 []int
fmt.Println(FromSlice(s1).ReduceP(2, 0, sum))
fmt.Println(FromSlice(s2).ReduceP(2, 0, sum))
// Output:
// 15
// 0
```
### ForEachP
**ForEachP()** performs an action for each item in a sequence.  
This is parallel version of **ForEach()** using specified number of goroutines to process items in parallel.
```go
s := []string{"Huey", "Dewey", "Louie"}
FromSlice(s).ForEachP(2, func(name string) {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    fmt.Println(name)
})
// Unordered output:
// Huey
// Dewey
// Louie
```
### ForEnumeratedP
**ForEnumeratedP()** is similar to **ForEachP()** but passes item index as first parameter of action function.  
This is parallel version of **ForEnumerated()** using specified number of goroutines to process items in parallel.
```go
s := []string{"Huey", "Dewey", "Louie"}
FromSlice(s).ForEnumeratedP(2, func(i int, name string) {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    fmt.Println(i, name)
})
// Unordered output:
// 0 Huey
// 1 Dewey
// 2 Louie
```
### AllP
**AllP()** determines whether all items of a sequence satisfy the given condition.  
This is parallel version of **All()** using specified number of goroutines to process items in parallel.
```go
isEven := func(i int) bool {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return i%2 == 0
}
s1 := []int{2, 4, 6, 8, 10, 12, 14, 16}
s2 := []int{2, 4, 6, 8, 11, 12, 14, 16}
fmt.Println(FromSlice(s1).AllP(2, isEven))
fmt.Println(FromSlice(s2).AllP(2, isEven))
// Output:
// true
// false
```
### AnyP
**AnyP()** determines whether any item of a sequence satisfies the given condition.  
This is parallel version of **Any()** using specified number of goroutines to process items in parallel.
```go
isEven := func(i int) bool {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return i%2 == 0
}
s1 := []int{1, 3, 5, 7, 9, 11, 13, 15}
s2 := []int{1, 2, 4, 7, 9, 11, 13, 16}
fmt.Println(FromSlice(s1).AnyP(2, isEven))
fmt.Println(FromSlice(s2).AnyP(2, isEven))
// Output:
// false
// true
```
### PartitionP
**PartitionP()** returns two slices. The first one contains items satisfying the given condition, the second one contains the rest of them.  
This is parallel version of **Partition()** using specified number of goroutines to process items in parallel.
```go
s := FromSlice([]int{1, 2, 3, 4, 5, 6})
t, f := s.PartitionP(1, func(i int) bool {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return i%2 == 0
})
for _, i := range t {
    fmt.Println("even:", i)
}
for _, i := range f {
    fmt.Println("odd:", i)
}
// Unordered output:
// even: 2
// even: 4
// even: 6
// odd: 1
// odd: 3
// odd: 5
```
### ZipP
**ZipP()** combines items of two sequences pairwise using a projection function.  
This is parallel version of **Zip()** using specified number of goroutines to process items in parallel.
```go
s1 := FromSlice([]int{1, 2, 3})
s2 := FromSlice([]int{4, 5, 6, 7})
ZipP(2, s1, s2, func(u int, v int) [2]int {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return [2]int{u, v}
}).ForEach(func(s [2]int) {
    fmt.Println(s)
})
// Unordered output:
// [1 4]
// [2 5]
// [3 6]
```
### Zip3P
**Zip3P()** is same as **ZipP()** but combines three sequences.  
This is parallel version of **Zip3()** using specified number of goroutines to process items in parallel.
```go
s1 := FromSlice([]int{1, 2, 3})
s2 := FromSlice([]int{4, 5, 6})
s3 := FromSlice([]int{7, 8, 9, 10})
Zip3P(2, s1, s2, s3, func(u int, v int, w int) [3]int {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    return [3]int{u, v, w}
}).ForEach(func(s [3]int) {
    fmt.Println(s)
})
// Unordered output:
// [1 4 7]
// [2 5 8]
// [3 6 9]
```