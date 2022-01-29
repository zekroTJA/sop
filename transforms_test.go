package sop

import (
	"fmt"
	"strconv"
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

func TestGroup(t *testing.T) {
	w := Slice([]int{1, 2, 3, 2})
	m := Group[int](w, func(v, i int) (mk string, mv int) {
		mk = strconv.Itoa(v)
		mv = v
		return
	})
	assert.Equal(t, map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}, m)

	assert.Panics(t, func() {
		Group[int, string, int](Slice([]int{1}), nil)
	})
}

func TestGroupE(t *testing.T) {
	type obj struct {
		K string
		V int
	}

	s := Slice([]obj{
		{"a", 1},
		{"a", 2},
		{"a", 3},
		{"b", 1},
		{"b", 2},
		{"c", 1},
	})

	m := GroupE[obj](s, func(v obj, i int) (string, int) {
		return v.K, v.V
	})

	assert.Equal(t, map[string]Enumerable[int]{
		"a": &slice[int]{[]int{1, 2, 3}},
		"b": &slice[int]{[]int{1, 2}},
		"c": &slice[int]{[]int{1}},
	}, m)
}

func TestMapFlat(t *testing.T) {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	s := MapFlat(m)
	st := []Tuple[string, int]{
		{"1", 1},
		{"2", 2},
		{"3", 3},
	}
	assert.ElementsMatch(t, st, s.Unwrap())
}
