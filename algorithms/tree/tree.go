package tree

type Node struct {
	Value int32
	Left  *Node
	Right *Node
}

func Equals(n1 *Node, n2 *Node) bool {
	if n1 == nil || n2 == nil {
		return n1 == n2
	}
	return n1.Value == n2.Value && Equals(n1.Left, n2.Left) && Equals(n1.Right, n2.Right)
}

func PreOrderTraverse(n *Node, f func(n *Node)) {
	f(n)
	if n.Left != nil {
		PreOrderTraverse(n.Left, f)
	}
	if n.Right != nil {
		PreOrderTraverse(n.Right, f)
	}
}

func InOrderTraverse(n *Node, f func(n *Node)) {
	if n.Left != nil {
		InOrderTraverse(n.Left, f)
	}
	f(n)
	if n.Right != nil {
		InOrderTraverse(n.Right, f)
	}
}

func PostOrderTraverse(n *Node, f func(n *Node)) {
	if n.Left != nil {
		PostOrderTraverse(n.Left, f)
	}
	if n.Right != nil {
		PostOrderTraverse(n.Right, f)
	}
	f(n)
}

func BreadthFirstSearch(n *Node, f func(n *Node)) {
	queue := []*Node{}
	queue = append(queue, n)
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		f(next)
		if next.Left != nil {
			queue = append(queue, next.Left)
		}
		if next.Right != nil {
			queue = append(queue, next.Right)
		}
	}
}
