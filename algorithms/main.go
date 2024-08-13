package main

import (
	"algorithms/tree"
	"fmt"
)

func createTree() *tree.Node {
	lln := tree.Node{
		Value: 3,
	}
	lrn := tree.Node{
		Value: 4,
	}
	ln := tree.Node{
		Value: 2,
		Left:  &lln,
		Right: &lrn,
	}
	rln := tree.Node{
		Value: 6,
	}
	rrn := tree.Node{
		Value: 7,
	}
	rn := tree.Node{
		Value: 5,
		Left:  &rln,
		Right: &rrn,
	}
	root := tree.Node{
		Value: 1,
		Left:  &ln,
		Right: &rn,
	}
	return &root
}




func main() {
	root := createTree()
	tree.BreadthFirstSearch(root, func(n *tree.Node) {
        fmt.Println(n.Value)
	})

	//CRYSTAL
	// fmt.Println("Hello, World!")
	// breaks := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	// fmt.Println(CystalBall(breaks))

	//BUBBLE
	// arr := []int{1, 3, 7, 4, 2}
	// fmt.Println(BubbleSort(arr))

	// LINKED LISTS
	// l := DoublyLinkedList[int]{nil, nil}
	// l.Append(1)
	// l.Append(2)
	// l.Append(3)
	// l.Append(4)
	// l.Append(3)
	// l.Remove(3)
	// l.Display()

	// q := Queue[int]{nil, nil}
	// q.Enqueue(1)
	// q.Enqueue(2)
	// q.Enqueue(3)
	// q.Enqueue(4)
	// fmt.Println(q.Peek())
	// fmt.Println(q.Deque())
	// fmt.Println(q.Deque())
	// fmt.Println(q.Deque())
	// fmt.Println(q.Deque())
	// fmt.Println(q.Deque())

	// // STACK
	// s := Stack[int]{nil}
	// s.Push(1)
	// s.Push(2)
	// s.Push(3)
	// s.Push(4)
	// s.Push(5)
	// fmt.Println(s.Peek())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// maze := amazeing.CreateMaze()
    // amazeing.SolveMaze(&maze)
}
