package sop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	w := Wrap([]int{1, 2, 3})

	r := Map(w, func(v int, i int) string {
		return fmt.Sprintf("%d: %d", i, v)
	})
	assert.Equal(t, r.Unwrap(), []string{"0: 1", "1: 2", "2: 3"})

	assert.Panics(t, func() {
		Map[int, int](Wrap([]int{1}), nil)
	})
}

func TestFlat(t *testing.T) {
	w := Wrap([][]int{
		{1, 2, 3},
		{4, 5},
		{6},
		{},
		{7, 8, 9, 10},
	})

	wf := Flat(w)
	assert.Equal(t, wf, Wrap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
