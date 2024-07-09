package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePolishNotation(t *testing.T) {

    r1 := ReversePolishNotation([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
    assert.Equal(t, 22, r1)
}
