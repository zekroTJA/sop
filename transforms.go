package sop

import "constraints"

// Map takes a Slice s and performs the passed function f
// on each element of the Slice s. The return value of the
// function f for each element is then packed into a new
// Slice in the same order as Slice s.
//
// f is getting passed the value v at the given position
// in the slice as well as the current index i.
func Map[TIn, TOut any](s Enumerable[TIn], f func(v TIn, i int) TOut) Enumerable[TOut] {
	notNil("f", f)
	res := newSliceFrom[TIn, TOut](s)
	s.Each(func(v TIn, i int) {
		res[i] = f(v, i)
	})
	return Slice(res)
}

// Flat takes a slice containing arrays and creates a
// new slice with all elements of the sub-arrays
// arranged into a one-dimensional array.
func Flat[T any](s Enumerable[[]T]) (res Enumerable[T]) {
	var i int
	s.Each(func(v []T, _ int) {
		i += len(v)
	})
	r := &slice[T]{make([]T, i)}
	i = 0
	s.Each(func(v []T, _ int) {
		for _, uv := range v {
			r.s[i] = uv
			i++
		}
	})
	res = r
	return
}

// Fill creates an empty Slice[T] with the given
// size n and executes f for each element in the
// Slice and sets the value at the given position
// to its return value.
//
// f is therefore getting passed the current
// index i in the slice.
func Fill[T any](n int, f func(i int) T) (res Enumerable[T]) {
	notNil("f", f)
	r := Slice(make([]T, n))
	for i := 0; i < n; i++ {
		r.s[i] = f(i)
	}
	res = r
	return
}

// Range creates an integer Slice filled with
// sequential numbers starting with s and ending
// with s+n-1 [n, s+n).
func Range[T constraints.Integer](s, n T) (res Enumerable[T]) {
	r := Slice(make([]T, n))
	for i := s; i < s+n; i++ {
		r.s[i-s] = i
	}
	res = r
	return
}
