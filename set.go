package sop

// set wraos a slice and ensures that
// each element in the set is unique
// inside the set.
type set[T comparable] struct {
	*slice[T]
}

var _ Enumerable[any] = (*set[any])(nil)

// Set packs a given slice []T into a
// *set[T] object. A set acts as same as
// a slice but guarantees that each element
// in the set is unique inside the set.
func Set[T comparable](s []T) (st *set[T]) {
	st = &set[T]{
		slice: &slice[T]{},
	}
	st.Append(Slice(s))
	return
}

// Contains returns true if the given
// element v is contained in the set.
func (s *set[T]) Contains(v T) bool {
	return s.Any(func(c T, _ int) bool {
		return c == v
	})
}

func (s *set[T]) Push(v T) {
	if !s.Contains(v) {
		s.slice.Push(v)
	}
}

func (s *set[T]) Append(v Enumerable[T]) {
	v.Each(func(e T, _ int) {
		s.Push(e)
	})
}
