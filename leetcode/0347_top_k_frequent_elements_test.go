package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {
	r1 := TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	assert.Equal(t, []int{1, 2}, r1)
}
