package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
    m := MinStackConstructor()
    m.Push(3)
    minV := m.GetMin()
    top := m.Top()
    assert.Equal(t, 3, minV)
    assert.Equal(t, 3, top)

    m.Push(4)
    minV = m.GetMin()
    top = m.Top()
    assert.Equal(t, 3, minV)
    assert.Equal(t, 4, top)


    m.Push(1)
    minV = m.GetMin()
    top = m.Top()
    assert.Equal(t, 1, minV)
    assert.Equal(t, 1, top)

    m.Pop()
    minV = m.GetMin()
    top = m.Top()
    assert.Equal(t, 3, minV)
    assert.Equal(t, 4, top)
}
