package heap_test

import (
	"algorithms/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestHeap(t *testing.T){
    h := heap.NewHeap()
    h.Insert(2)
    h.Insert(5)
    h.Insert(10)
    h.Insert(3)
    h.Insert(2)
    h.Insert(4)
    h.Insert(8)
    h.Insert(1)
    assert.Equal(t, []int{1, 2, 4, 2, 3, 10, 8, 5}, h.GetArr())

    h.DeleteI(0)
    assert.Equal(t, []int{2, 2, 4, 5, 3, 10, 8}, h.GetArr())
}


