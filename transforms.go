package sop

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

// Flat takes a slice containing arrays and creates a
// new slice with all elements of the sub-arrays
// arranged into a one-dimensional array.
func Flat[T any](s Slice[[]T]) (res Slice[T]) {
	var i int
	s.Each(func(v []T, _ int) {
		i += len(v)
	})
	res = Slice[T]{make([]T, i)}
	i = 0
	s.Each(func(v []T, _ int) {
		for _, uv := range v {
			res.s[i] = uv
			i++
		}
	})
	return
}
