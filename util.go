package sop

import (
	"fmt"
	"reflect"
)

func notNil(name string, v interface{}) {
	if v == nil || (reflect.ValueOf(v).IsNil()) {
		panic(fmt.Sprintf("parameter %s can not be nil", v))
	}
}

func newSliceFrom[TIn, TOut any](s Enumerable[TIn]) []TOut {
	return make([]TOut, s.Len())
}

func copySlice[T any](t []T) (s []T) {
	s = make([]T, len(t))
	copy(s, t)
	return
}
