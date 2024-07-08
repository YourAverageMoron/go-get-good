package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinnerCircularGame(t *testing.T) {
	r1 := WinnerCircularGame(5, 2)
	assert.Equal(t, 3, r1)
}
