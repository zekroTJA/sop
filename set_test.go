package sop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := Set([]int{1, 2, 3, 4, 1, 3, 5, 6, 2})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, s.Unwrap())

	s = Set(([]int)(nil))
	assert.Equal(t, []int(nil), s.Unwrap())
}

func TestSetPush(t *testing.T) {
	s := Set([]int{2, 3})

	s.Push(1)
	assert.Equal(t, []int{2, 3, 1}, s.Unwrap())

	s.Push(2)
	assert.Equal(t, []int{2, 3, 1}, s.Unwrap())
}

func TestSetAppend(t *testing.T) {
	s := Set([]int{2, 3})

	s.Append(Slice([]int{1, 4, 5}))
	assert.Equal(t, []int{2, 3, 1, 4, 5}, s.Unwrap())

	s.Append(Slice([]int{7, 1, 9, 4, 5}))
	assert.Equal(t, []int{2, 3, 1, 4, 5, 7, 9}, s.Unwrap())
}
