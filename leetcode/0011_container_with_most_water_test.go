package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {

    r1 := ContainerWithMostWater([]int{2,3,4,5,18,17,6})
    assert.Equal(t, 17, r1)
}
 
