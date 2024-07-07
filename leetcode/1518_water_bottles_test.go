package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaterBottles(t *testing.T) {
    r1 := WaterBottles(9,3)
	assert.Equal(t, 13, r1)
}
