package main

type ReverseLinkedListNode struct {
	Val  int
	Next *ReverseLinkedListNode
}

func ReverseLinkedList(head *ReverseLinkedListNode) *ReverseLinkedListNode {
	var prev *ReverseLinkedListNode
	for {
		if head == nil {
			return prev
		}
		head, head.Next, prev = head.Next, prev, head
	}
}
