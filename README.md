# sop &nbsp; [![](https://img.shields.io/badge/docs-pkg.do.dev-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/zekrotja/sop?tab=doc) [![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/zekrotja/sop)](https://github.com/zekroTJA/sop/releases) [![Unit Tests](https://github.com/zekroTJA/sop/actions/workflows/unittests.yml/badge.svg)](https://github.com/zekroTJA/sop/actions/workflows/unittests.yml)

✨ **W.I.P.** ✨

sop (`slices operation`) is a go1.18+ package to *(maybe)* simplify performing operations on slices in a fluent-like style with common operations like `map`, `filter`, `each`, `aggregate` and much more.

Currently, the following functionalities are covered:
```go
// Map takes a Slice s and performs the passed function f
// on each element of the Slice s. The return value of the
// function f for each element is then packed into a new
// Slice in the same order as Slice s.
//
// f is getting passed the value v at the given position
// in the slice as well as the current index i.
func Map[TIn, TOut any](s Enumerable[TIn], f func(v TIn, i int) TOut) Enumerable[TOut]

// Fill creates an empty Slice[T] with the given
// size n and executes f for each element in the
// Slice and sets the value at the given position
// to its return value.
//
// f is therefore getting passed the current
// index i in the slice.
func Fill[T any](n int, f func(i int) T) (res Enumerable[T])

// Flat takes a slice containing arrays and creates a
// new slice with all elements of the sub-arrays
// arranged into a one-dimensional array.
func Flat[T any](s Enumerable[[]T]) (res Enumerable[T])

// Range creates an integer Slice filled with
// sequential numbers starting with s and ending
// with s+n-1 [n, s+n).
func Range[T constraints.Integer](s, n T) (res Enumerable[T])

// Enumerable specifies a wrapped slice object
// to perform different enumerable operations
// on.
type Enumerable[T any] interface {
	// Unwrap returns the originaly packed
	// Enumerable []T of the Enumerable[T] object.
	Unwrap() []T
	// Len returns the length of the given Enumerable.
	Len() int
	// Each performs the given function f on each
	// element in the Enumerable.
	//
	// f is getting passed the value v at the
	// current position as well as the current
	// index i.
	Each(f func(v T, i int))
	// Filter performs preticate p on each element
	// in the Enumerable and each element where p
	// returns true will be added to the result
	// Enumerable.
	//
	// p is getting passed the value v at the
	// current position as well as the current
	// index i.
	Filter(p func(v T, i int) bool) Enumerable[T]
	// Any returns true when at least one element in
	// the given Enumerable result in a true return of p
	// when performed on p.
	Any(p func(v T, i int) bool) bool
	// All returns true when all elements in the given
	// Enumerable result in a true return of p when performed
	// on p.
	All(p func(v T, i int) bool) bool
	// All returns true when all elements in the given
	// Enumerable result in a true return of p when performed
	// on p.
	None(p func(v T, i int) bool) bool
	// First returns the value and index of the first
	// occurence in the Enumerable where preticate p
	// returns true.
	//
	// If this applies to no element in the Enumerable,
	// default of T and -1 is returned.
	First(p func(v T, i int) bool) (T, int)
	// Count returns the number of elements in the given
	// Enumerable which, when applied on p, return true.
	Count(p func(v T, i int) bool) int
	// Shuffle re-arranges the given Enumerable in a pseudo-random
	// order and returns the result Enumerable.
	//
	// You can also pass a custom random source rngSrc if
	// you desire.
	Shuffle(rngSrc ...rand.Source) Enumerable[T]
	// Sort re-orders the Enumerable x given the provided less function.
	Sort(less func(p, q T, i int) bool) Enumerable[T]
	// Aggregate applies tze multiplicator function f over
	// all elements of the given Enumerable and returns the final
	// result.
	Aggregate(f func(a, b T) T) T
	// Push appends the passed value v to the Enumerable.
	Push(v T)
	// Pop removes the last element of the Enumerable and
	// returns its value. If the Enumerable is empty, the
	// default value of T is returned.
	Pop() T
	// Append adds all elements of Enumerable v to the
	// end of the current Enumerable.
	Append(v Enumerable[T])
	// Flush removes all elements of the given Enumerable.
	Flush()
	// Enumerable removes the values from the given Enumerable
	// starting at i with the amount of n. The removed
	// Enumerable is returned as new Enumerable.
	Splice(i, n int) Enumerable[T]
	// At safely accesses the element in the Enumerable
	// at the given index i and returns it, if existent.
	// If there is no value at i, default of T and false
	// is returned.
	At(i int) (T, bool)
	// Replace safely replaces the value in the Enumerable
	// at the given index i with the given value v and
	// returns true if the value was replaced. If the
	// Enumerable has no value at i, false is returned.
	Replace(i int, v T) bool
}
```