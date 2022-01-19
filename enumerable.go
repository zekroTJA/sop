package sop

import "math/rand"

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
