package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	r1 := ValidParentheses("{}[](){")
	assert.Equal(t, false, r1)
	r2 := ValidParentheses("{}[]()}")
	assert.Equal(t, false, r2)
	r3 := ValidParentheses("{}[]()")
	assert.Equal(t, true, r3)
}
