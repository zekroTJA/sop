package sop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	var w Enumerable[int] = Slice([]int{1, 2, 3})

	r := Map(w, func(v int, i int) string {
		return fmt.Sprintf("%d: %d", i, v)
	})
	assert.Equal(t, []string{"0: 1", "1: 2", "2: 3"}, r.Unwrap())

	assert.Panics(t, func() {
		Map[int, int](Slice([]int{1}), nil)
	})
}

func TestFlat(t *testing.T) {
	var w Enumerable[[]int] = Slice([][]int{
		{1, 2, 3},
		{4, 5},
		{6},
		{},
		{7, 8, 9, 10},
	})

	wf := Flat(w)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, wf.Unwrap())
}

func TestFill(t *testing.T) {
	w := Fill(5, func(i int) int {
		return i + 1
	})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, w.Unwrap())
}

func TestRange(t *testing.T) {
	w := Range(3, 5)
	assert.Equal(t, []int{3, 4, 5, 6, 7}, w.Unwrap())
}
