package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
    r1:= ValidPalindrome("A man, a plan, a canal: Panama")
	assert.True(t, r1)

    r2:= ValidPalindrome("race a car")
	assert.True(t, !r2)

    r3:= ValidPalindrome(" ")
	assert.True(t, r3)
}
