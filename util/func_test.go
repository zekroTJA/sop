package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zekrotja/sop"
)

func TestAsc(t *testing.T) {
	s := sop.Slice([]int{2, 4, 5, 1, 3})
	r := s.Sort(Asc[int])
	assert.Equal(t, []int{1, 2, 3, 4, 5}, r.Unwrap())
}

func TestDesc(t *testing.T) {
	s := sop.Slice([]int{2, 4, 5, 1, 3})
	r := s.Sort(Desc[int])
	assert.Equal(t, []int{5, 4, 3, 2, 1}, r.Unwrap())
}

func TestNotNil(t *testing.T) {
	type a struct{}
	s := sop.Slice([]*a{{}, nil})
	r := s.Filter(NotNil[a])
	assert.Equal(t, []*a{{}}, r.Unwrap())
}
