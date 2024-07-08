package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrappingRainWater(t *testing.T) {
    r1:= TrappingRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
    assert.Equal(t, 6, r1)
} 
