package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarFleet(t *testing.T) {
    r1:=CarFleet(10, []int{6,8}, []int{3,2})
    assert.Equal(t, 2, r1) 
} 
