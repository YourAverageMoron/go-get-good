package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	r1 := LongestConsecutiveSequence([]int{100, 4, 200, 1, 3, 2})
    assert.Equal(t, 4, r1)

	r2 := LongestConsecutiveSequence([]int{2,3,4,6,7,8,5,1})
    assert.Equal(t, 8, r2)
}
