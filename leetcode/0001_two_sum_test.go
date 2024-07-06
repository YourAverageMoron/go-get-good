package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumLists(t *testing.T) {
    r1 := TwoSumLists([]int{1, 2, 3, 4}, 7)
    assert.Equal(t, []int{3,2}, r1)
}


func TestTwoSumMap(t *testing.T) {
    r1 := TwoSumMap([]int{1, 2, 3, 4}, 7)
    assert.Equal(t, []int{3,2}, r1)
}
	
