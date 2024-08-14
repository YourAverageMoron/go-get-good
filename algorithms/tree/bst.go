package tree

func BstInsert(n *Node, root *Node) {
	if root == nil {
		root = n
		return
	}
	if n.Value <= root.Value {
		if root.Left == nil {
			root.Left = n
			return
		}
		BstInsert(n, root.Left)
		return
	}
	if root.Right == nil {
		root.Right = n
		return
	}
	BstInsert(n, root.Right)
}

func BstSearch(value int32, n *Node) bool {
	if n.Value == value {
		return true
	}
	if value <= n.Value && n.Left != nil {
		return BstSearch(value, n.Left)
	}
	if n.Right != nil {
		return BstSearch(value, n.Right)
	}
	return false
}
