package main


import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestReverseSubstringsBetweenEachPairOfParentheses(t *testing.T) {
    r1 := ReverseSubstringsBetweenEachPairOfParentheses("q(abc)z")
    assert.Equal(t, "qcbaz", r1)
} 
