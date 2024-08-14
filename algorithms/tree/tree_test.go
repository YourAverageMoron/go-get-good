package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTree() *Node {
	lln := Node{
		Value: 3,
	}
	lrn := Node{
		Value: 4,
	}
	ln := Node{
		Value: 2,
		Left:  &lln,
		Right: &lrn,
	}
	rln := Node{
		Value: 6,
	}
	rrn := Node{
		Value: 7,
	}
	rn := Node{
		Value: 5,
		Left:  &rln,
		Right: &rrn,
	}
	root := Node{
		Value: 1,
		Left:  &ln,
		Right: &rn,
	}
	return &root
}

func TestPreOrderTraverse(t *testing.T) {
	root := createTree()
	arr := []int32{}
	PreOrderTraverse(root, func(n *Node) {
		arr = append(arr, n.Value)
	})
	assert.Equal(t, []int32{1, 2, 3, 4, 5, 6, 7}, arr)
}

func TestInOrderTraverse(t *testing.T) {
	root := createTree()
	arr := []int32{}
	InOrderTraverse(root, func(n *Node) {
		arr = append(arr, n.Value)
	})
	assert.Equal(t, []int32{3, 2, 4, 1, 6, 5, 7}, arr)
}

func TestPostOrderTraverse(t *testing.T) {
	root := createTree()
	arr := []int32{}
	PostOrderTraverse(root, func(n *Node) {
		arr = append(arr, n.Value)
	})
	assert.Equal(t, []int32{3, 4, 2, 6, 7, 5, 1}, arr)
}

func TestBFSTraverse(t *testing.T) {
	root := createTree()
	arr := []int32{}
	BreadthFirstSearch(root, func(n *Node) {
		arr = append(arr, n.Value)
	})
	assert.Equal(t, []int32{1, 2, 5, 3, 4, 6, 7}, arr)
}

func TestEquals(t *testing.T) {
	root := createTree()
	root2 := createTree()
	equal := Equals(root, root2)
	assert.Equal(t, true, equal)
    notEqual := Equals(root, root2.Left)
	assert.Equal(t, false, notEqual)
}
