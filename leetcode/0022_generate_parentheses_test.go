package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
    r1 := GenerateParentheses(1)
	assert.Equal(t, []string{"()"}, r1)
}
