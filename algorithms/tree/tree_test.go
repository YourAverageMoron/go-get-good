package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTree() *Node {
	lln := Node{
		value: 3,
	}
	lrn := Node{
		value: 4,
	}
	ln := Node{
		value: 2,
		left:  &lln,
		right: &lrn,
	}
	rln := Node{
		value: 6,
	}
	rrn := Node{
		value: 7,
	}
	rn := Node{
		value: 5,
		left:  &rln,
		right: &rrn,
	}
	root := Node{
		value: 1,
		left:  &ln,
		right: &rn,
	}
	return &root
}

func TestPreOrderTraverse(t *testing.T) {
	root := createTree()
	arr := []int32{}
	PreOrderTraverse(root, func(n *Node) {
		arr = append(arr, n.value)
	})
	assert.Equal(t, []int32{1, 2, 3, 4, 5, 6, 7}, arr)
}

func TestInOrderTraverse(t *testing.T) {
	root := createTree()
	arr := []int32{}
	InOrderTraverse(root, func(n *Node) {
		arr = append(arr, n.value)
	})
	assert.Equal(t, []int32{3, 2, 4, 1, 6, 5, 7}, arr)
}

func TestPostOrderTraverse(t *testing.T) {
	root := createTree()
	arr := []int32{}
	PostOrderTraverse(root, func(n *Node) {
		arr = append(arr, n.value)
	})
	assert.Equal(t, []int32{3, 4, 2, 6, 7, 5, 1}, arr)
}
