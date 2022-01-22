package util

import "constraints"

// Asc is a function used to sort
// enumerables in ascending order.
func Asc[T constraints.Ordered](p, q T, i int) bool {
	return p < q
}

// Desc is a function used to sort
// enumerables in descending order.
func Desc[T constraints.Ordered](p, q T, i int) bool {
	return p > q
}

// NolNil is a preticate function
// checking if the element reference
// is not nil.
func NotNil[T any](v *T, i int) bool {
	return v != nil
}
