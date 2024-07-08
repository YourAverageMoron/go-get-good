package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
    r1 := ThreeSumTwoPointer([]int{5,5,5})
    assert.Equal(t, [][]int{}, r1)
}
 
