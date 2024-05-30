package tree

import "fmt"

type Node struct {
	value  int32
	left   *Node
	right  *Node
	parent *Node
}





func PreOrderTraverse(n *Node) {
	fmt.Println(n.value)
	if n.left != nil {
		PreOrderTraverse(n.left)
	}
	if n.right != nil {
		PreOrderTraverse(n.right)
	}
}
