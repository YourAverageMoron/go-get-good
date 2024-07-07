package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {

    r1 := ProductOfArrayExceptSelf([]int{1,2,3,4})
	assert.Equal(t, []int{24,12,8,6}, r1)
}
