package sop

import (
	"math/rand"
	"sort"
	"time"
)

// slice wraps a native slice to perform
// operations on it.
type slice[T any] struct {
	s []T
}

var _ Enumerable[any] = (*slice[any])(nil)

// Slice packs a given slice []T into a
// *slice[T] object.
func Slice[T any](s []T) *slice[T] {
	return &slice[T]{s}
}

// Unwrap returns the originaly packed
// slice []T of the Slice[T] object.
func (s *slice[T]) Unwrap() []T {
	return s.s
}

// Len returns the length of the given Slice.
func (s *slice[T]) Len() int {
	return len(s.s)
}

// Each performs the given function f on each
// element in the Slice.
//
// f is getting passed the value v at the
// current position as well as the current
// index i.
func (s *slice[T]) Each(f func(v T, i int)) {
	notNil("f", f)
	for i, v := range s.s {
		f(v, i)
	}
}

// Filter performs preticate p on each element
// in the Slice and each element where p
// returns true will be added to the result
// slice.
//
// p is getting passed the value v at the
// current position as well as the current
// index i.
func (s *slice[T]) Filter(p func(v T, i int) bool) Enumerable[T] {
	notNil("p", p)
	res := newSliceFrom[T, T](s)
	var j int
	s.Each(func(v T, i int) {
		if p(v, i) {
			res[j] = v
			j++
		}
	})
	return Slice(res[:j])
}

// Any returns true when at least one element in
// the given slice result in a true return of p
// when performed on p.
func (s *slice[T]) Any(p func(v T, i int) bool) bool {
	notNil("p", p)
	for i, v := range s.s {
		if p(v, i) {
			return true
		}
	}
	return false
}

// All returns true when all elements in the given
// slice result in a true return of p when performed
// on p.
func (s *slice[T]) All(p func(v T, i int) bool) bool {
	notNil("p", p)
	return !s.Any(func(v T, i int) bool {
		return !p(v, i)
	})
}

// All returns true when all elements in the given
// slice result in a true return of p when performed
// on p.
func (s *slice[T]) None(p func(v T, i int) bool) bool {
	notNil("p", p)
	return !s.Any(p)
}

// First returns the value and index of the first
// occurence in the Enumerable where preticate p
// returns true.
//
// If this applies to no element in the Enumerable,
// default of T and -1 is returned.
func (s *slice[T]) First(p func(v T, i int) bool) (rv T, ri int) {
	ri = -1
	s.Any(func(v T, i int) bool {
		ok := p(v, i)
		if ok {
			rv = v
			ri = i
		}
		return ok
	})
	return
}

// Count returns the number of elements in the given
// slice which, when applied on p, return true.
func (s *slice[T]) Count(p func(v T, i int) bool) (c int) {
	notNil("p", p)
	s.Each(func(v T, i int) {
		if p(v, i) {
			c++
		}
	})
	return
}

// Shuffle re-arranges the given slice in a pseudo-random
// order and returns the result slice.
//
// You can also pass a custom random source rngSrc if
// you desire.
func (s *slice[T]) Shuffle(rngSrc ...rand.Source) (res Enumerable[T]) {
	var rng *rand.Rand
	if len(rngSrc) != 0 {
		rng = rand.New(rngSrc[0])
	} else {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	src := copySlice(s.s)
	r := Slice(newSliceFrom[T, T](s))
	for i := len(src) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		r.s[i] = src[j]
		src = append(src[:j], src[j+1:]...)
	}
	r.s[0] = src[0]
	res = r
	return
}

// Sort re-orders the slice x given the provided less function.
func (s *slice[T]) Sort(less func(p, q T, i int) bool) Enumerable[T] {
	notNil("less", less)
	res := copySlice(s.s)
	sort.Slice(res, func(i, j int) bool {
		return less(res[i], res[j], i)
	})
	return Slice(res)
}

// Aggregate applies tze multiplicator function f over
// all elements of the given Slice and returns the final
// result.
func (s *slice[T]) Aggregate(f func(a, b T) T) (c T) {
	notNil("f", f)
	if s.Len() == 0 {
		return
	}
	c = s.s[0]
	for i := 1; i < s.Len(); i++ {
		c = f(c, s.s[i])
	}
	return
}

// Push appends the passed value v to the Slice.
func (s *slice[T]) Push(v T) {
	s.s = append(s.s, v)
}

// Pop removes the last element of the Slice and
// returns its value. If the Slice is empty, the
// default value of T is returned.
func (s *slice[T]) Pop() (res T) {
	if s.Len() == 0 {
		return
	}
	res = s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return
}

// Append adds all elements of Slice v to the
// end of the current slice.
func (s *slice[T]) Append(v Enumerable[T]) {
	s.s = append(s.s, v.Unwrap()...)
}

// Flush removes all elements of the given Slice.
func (s *slice[T]) Flush() {
	s.s = make([]T, 0)
}

// Slice removes the values from the given slice
// starting at i with the amount of n. The removed
// slice is returned as new Slice.
func (s *slice[T]) Splice(i, n int) (res Enumerable[T]) {
	t := copySlice(s.s)
	res = Slice(t[i : i+n])
	s.s = append(s.s[:i], s.s[i+n:]...)
	return
}

// At safely accesses the element in the Enumerable
// at the given index i and returns it, if existent.
// If there is no value at i, default of T and false
// is returned.
func (s *slice[T]) At(i int) (v T, ok bool) {
	if i < 0 || i >= s.Len() {
		return
	}
	v = s.s[i]
	ok = true
	return
}

// Replace safely replaces the value in the Enumerable
// at the given index i with the given value v and
// returns true if the value was replaced. If the
// Enumerable has no value at i, false is returned.
func (s *slice[T]) Replace(i int, v T) (ok bool) {
	if i < 0 || i >= s.Len() {
		return
	}
	s.s[i] = v
	ok = true
	return
}
