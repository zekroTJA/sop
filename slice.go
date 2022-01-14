package sop

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

// Slice wraps a native slice to perform
// operations on it.
type Slice[T any] struct {
	s []T
}

// Wrap packs a given slice []T into a
// Slice[T] object.
func Wrap[T any](s []T) Slice[T] {
	return Slice[T]{s}
}

// Map takes a Slice s and performs the passed function f
// on each element of the Slice s. The return value of the
// function f for each element is then packed into a new
// Slice in the same order as Slice s.
//
// f is getting passed the value v at the given position
// in the slice as well as the current index i.
func Map[TIn, TOut any](s Slice[TIn], f func(v TIn, i int) TOut) Slice[TOut] {
	notNil("f", f)
	res := newSliceFrom[TIn, TOut](s)
	s.Each(func(v TIn, i int) {
		res[i] = f(v, i)
	})
	return Wrap(res)
}

// Unwrap returns the originaly packed
// slice []T of the Slice[T] object.
func (s Slice[T]) Unwrap() []T {
	return s.s
}

// Len returns the length of the given Slice.
func (s Slice[T]) Len() int {
	return len(s.s)
}

// Each performs the given function f on each
// element in the Slice.
//
// f is getting passed the value v at the
// current position as well as the current
// index i.
func (s Slice[T]) Each(f func(v T, i int)) {
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
func (s Slice[T]) Filter(p func(v T, i int) bool) Slice[T] {
	notNil("p", p)
	res := newSliceFrom[T, T](s)
	var j int
	s.Each(func(v T, i int) {
		if p(v, i) {
			res[j] = v
			j++
		}
	})
	return Wrap(res[:j])
}

// Any returns true when at least one element in
// the given slice result in a true return of p
// when performed on p.
func (s Slice[T]) Any(p func(v T, i int) bool) bool {
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
func (s Slice[T]) All(p func(v T, i int) bool) bool {
	notNil("p", p)
	return !s.Any(func(v T, i int) bool {
		return !p(v, i)
	})
}

// All returns true when all elements in the given
// slice result in a true return of p when performed
// on p.
func (s Slice[T]) None(p func(v T, i int) bool) bool {
	notNil("p", p)
	return !s.Any(p)
}

// Count returns the number of elements in the given
// slice which, when applied on p, return true.
func (s Slice[T]) Count(p func(v T, i int) bool) (c int) {
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
func (s Slice[T]) Shuffle(rngSrc ...rand.Source) (res Slice[T]) {
	var rng *rand.Rand
	if len(rngSrc) != 0 {
		rng = rand.New(rngSrc[0])
	} else {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	src := copySlice(s.s)
	res = Wrap(newSliceFrom[T, T](s))
	for i := len(src) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		res.s[i] = src[j]
		src = append(src[:j], src[j+1:]...)
	}
	res.s[0] = src[0]
	return
}

// Sort re-orders the slice x given the provided less function.
func (s Slice[T]) Sort(less func(p, q T, i int) bool) Slice[T] {
	notNil("less", less)
	res := copySlice(s.s)
	sort.Slice(res, func(i, j int) bool {
		return less(res[i], res[j], i)
	})
	return Wrap(res)
}

// Aggregate applies tze multiplicator function f over
// all elements of the given Slice and returns the final
// result.
func (s Slice[T]) Aggregate(f func(a, b T) T) (c T) {
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

func notNil(name string, v interface{}) {
	if v == nil || (reflect.ValueOf(v).IsNil()) {
		panic(fmt.Sprintf("parameter %s can not be nil", v))
	}
}

func newSliceFrom[TIn, TOut any](s Slice[TIn]) []TOut {
	return make([]TOut, s.Len())
}

func copySlice[T any](t []T) (s []T) {
	s = make([]T, len(t))
	copy(s, t)
	return
}
