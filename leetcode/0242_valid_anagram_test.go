package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	r1 := ValidAnagram("qwe", "ewq")
	assert.True(t, r1, "should return true valid anagram")

	r2 := ValidAnagram("qwe", "ewq1")
	assert.True(t, !r2, "should return false for strings of different length")

	r3 := ValidAnagram("qwe", "asd")
	assert.True(t, !r3, "should return false invalid anagrams")

}
