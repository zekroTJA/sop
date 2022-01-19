package sop

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	s := []int{1, 2, 3}
	w := Slice(s)
	assert.Equal(t, s, w.Unwrap())
}

func TestEach(t *testing.T) {
	w := Slice([]string{"1", "2", "3"})

	m := map[int]string{}
	w.Each(func(v string, i int) {
		m[i] = v
	})
	assert.Equal(t, m, map[int]string{
		0: "1",
		1: "2",
		2: "3",
	})

	assert.Panics(t, func() {
		Slice([]int{1}).Each(nil)
	})
}

func TestFilter(t *testing.T) {
	w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	r := w.Filter(func(v, _ int) bool {
		return v%2 == 0
	})
	assert.Equal(t, r.Unwrap(), []int{2, 4, 6, 8, 10})

	assert.Panics(t, func() {
		Slice([]int{1}).Filter(nil)
	})
}

func TestAny(t *testing.T) {
	w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res := w.Any(func(v, i int) bool {
		return v == 4
	})
	assert.True(t, res)

	res = w.Any(func(v, i int) bool {
		return v == 20
	})
	assert.False(t, res)

	assert.Panics(t, func() {
		Slice([]int{1}).Any(nil)
	})
}

func TestAll(t *testing.T) {
	w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res := w.All(func(v, i int) bool {
		return v < 11
	})
	assert.True(t, res)

	res = w.All(func(v, i int) bool {
		return v < 10
	})
	assert.False(t, res)

	assert.Panics(t, func() {
		Slice([]int{1}).All(nil)
	})
}

func TestNone(t *testing.T) {
	w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res := w.None(func(v, i int) bool {
		return v < 0
	})
	assert.True(t, res)

	res = w.None(func(v, i int) bool {
		return v < 10
	})
	assert.False(t, res)

	assert.Panics(t, func() {
		Slice([]int{1}).None(nil)
	})
}

func TestCount(t *testing.T) {
	w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res := w.Count(func(v, i int) bool {
		return v < 7
	})
	assert.Equal(t, res, 6)

	res = w.Count(func(v, i int) bool {
		return i == v
	})
	assert.Equal(t, res, 0)

	assert.Panics(t, func() {
		Slice([]int{1}).Count(nil)
	})
}

func TestShuffle(t *testing.T) {
	var w Enumerable[int] = Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < 1000; i++ {
		r := w.Shuffle()
		assert.NotEqual(t, w, r)
		w = r
	}

	w = Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t,
		w.Shuffle(rand.NewSource(1)),
		w.Shuffle(rand.NewSource(1)))
}

func TestSort(t *testing.T) {
	w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res := w.Shuffle()
	assert.NotEqual(t, w, res)

	srt := res.Sort(func(p, q, _ int) bool {
		return p < q
	})
	assert.Equal(t, w, srt)

	assert.Panics(t, func() {
		Slice([]int{1}).Sort(nil)
	})
}

func TestAggregate(t *testing.T) {
	{
		w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

		r := w.Aggregate(func(a, b int) int {
			return a + b
		})
		assert.Equal(t, 55, r)

		r = w.Aggregate(func(a, b int) int {
			return a * b
		})
		assert.Equal(t, 3_628_800, r)
	}

	{
		w := Slice([]string{"a", "b", "c"})

		r := w.Aggregate(func(a, b string) string {
			return a + b
		})
		assert.Equal(t, "abc", r)
	}

	{
		w := Slice([]int{})
		r := w.Aggregate(func(a, b int) int {
			return a + b
		})
		assert.Equal(t, 0, r)

		w = Slice[int](nil)
		r = w.Aggregate(func(a, b int) int {
			return a + b
		})
		assert.Equal(t, 0, r)
	}

	assert.Panics(t, func() {
		Slice([]int{1}).None(nil)
	})
}

func TestPush(t *testing.T) {
	w := Slice([]int{1, 2})
	w.Push(3)
	assert.Equal(t, []int{1, 2, 3}, w.Unwrap())
}

func TestPop(t *testing.T) {
	w := Slice([]int{1, 2, 3})
	assert.Equal(t, 3, w.Pop())
	assert.Equal(t, []int{1, 2}, w.Unwrap())

	w = Slice([]int{})
	assert.Equal(t, 0, w.Pop())
}

func TestAppend(t *testing.T) {
	w := Slice([]int{1, 2})
	w.Append(Slice([]int{3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, w.Unwrap())

	w = Slice([]int{1, 2})
	w.Append(Slice([]int{}))
	assert.Equal(t, []int{1, 2}, w.Unwrap())
}

func TestFlush(t *testing.T) {
	w := Slice([]int{1, 2})
	w.Flush()
	assert.Equal(t, []int{}, w.Unwrap())
}

func TestSplice(t *testing.T) {
	w := Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	r := w.Splice(4, 3)
	assert.Equal(t, []int{5, 6, 7}, r.Unwrap())
	assert.Equal(t, []int{1, 2, 3, 4, 8, 9, 10}, w.Unwrap())
}

func TestAt(t *testing.T) {
	w := Slice([]int{1, 2, 3})

	v, ok := w.At(-1)
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	v, ok = w.At(3)
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	v, ok = w.At(0)
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = w.At(2)
	assert.Equal(t, 3, v)
	assert.True(t, ok)
}

func TestReplace(t *testing.T) {
	w := Slice([]int{1, 2, 3})

	ok := w.Replace(-1, 4)
	assert.Equal(t, []int{1, 2, 3}, w.Unwrap())
	assert.False(t, ok)

	ok = w.Replace(3, 4)
	assert.Equal(t, []int{1, 2, 3}, w.Unwrap())
	assert.False(t, ok)

	ok = w.Replace(1, 4)
	assert.Equal(t, []int{1, 4, 3}, w.Unwrap())
	assert.True(t, ok)

	ok = w.Replace(2, 5)
	assert.Equal(t, []int{1, 4, 5}, w.Unwrap())
	assert.True(t, ok)
}
