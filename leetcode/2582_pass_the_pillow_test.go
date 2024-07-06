package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassThePillow(t *testing.T) {
    r1 := PassThePillowSolution(4, 5)
	assert.Equal(t, r1, 2)

    r2 := PassThePillowSolution(3, 2)
	assert.Equal(t, r2, 3)

    r3 := PassThePillowSolution(3, 10)
	assert.Equal(t, r3, 3)

    r4 := PassThePillowSolution(9, 4)
	assert.Equal(t, r4, 5)

    r5 := PassThePillowSolution(2, 5)
	assert.Equal(t, r5, 2)
}
