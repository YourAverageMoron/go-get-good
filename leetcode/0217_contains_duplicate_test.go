package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {
	r1 := ContainsDuplicate([]int{1, 2, 3})
	assert.True(t, !r1, "returns false when does not contain duplicate")

	r2 := ContainsDuplicate([]int{2, 2, 3})
	assert.True(t, r2, "returns true when contains duplicate")
}
