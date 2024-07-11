package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	n2 := &ReverseLinkedListNode{Val: 2, Next: nil}
	n1 := &ReverseLinkedListNode{Val: 1, Next: n2}

	r1 := ReverseLinkedList(n1)
	assert.Equal(t, n2, r1)
	assert.Equal(t, n1, r1.Next)
}
