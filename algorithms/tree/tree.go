package tree

type Node struct {
	value  int32
	left   *Node
	right  *Node
	parent *Node
}

func PreOrderTraverse(n *Node, f func(n *Node)) {
	f(n)
	if n.left != nil {
		PreOrderTraverse(n.left, f)
	}
	if n.right != nil {
		PreOrderTraverse(n.right, f)
	}
}


func InOrderTraverse(n *Node, f func(n *Node)) {
	if n.left != nil {
		InOrderTraverse(n.left, f)
	}
	f(n)
	if n.right != nil {
		InOrderTraverse(n.right, f)
	}
}

func PostOrderTraverse(n *Node, f func(n *Node)) {
	if n.left != nil {
		PostOrderTraverse(n.left, f)
	}
	if n.right != nil {
		PostOrderTraverse(n.right, f)
	}
	f(n)
}
