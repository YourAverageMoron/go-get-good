package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumIIInputArraySorted(t *testing.T) {
	r1 := TwoSumIIInputArraySorted([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, []int{1, 2}, r1)
}
