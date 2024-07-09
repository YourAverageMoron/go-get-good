package main


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
    r1:=DailyTemperatures([]int{73,74,75,71,69,72,76,73})
    assert.Equal(t, []int{1,1,4,2,1,1,0,0},r1)
} 
