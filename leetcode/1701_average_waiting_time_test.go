package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageWaitingTime(t *testing.T) {
	r1 := AverageWaitingTime([][]int{
		{1, 2}, {2, 5}, {4, 3},
	})
	assert.Equal(t, 5.0, r1)
}
