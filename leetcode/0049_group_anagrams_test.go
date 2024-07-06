package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	r1 := GroupAnagrams([]string{"ia", "qwe", "ewq", "ai"})
	assert.Equal(t, [][]string{{"ia", "ai"}, {"qwe", "ewq"}}, r1)
}
