package y2026

import "fmt"

type PrintTheElementsOfALinkedListSinglyLinkedListNode struct {
	data int32
	next *PrintTheElementsOfALinkedListSinglyLinkedListNode
}

func PrintTheElementsOfALinkedList(head *PrintTheElementsOfALinkedListSinglyLinkedListNode) {
    for head != nil {
        fmt.Println(head.data)
        head = head.next
    }
}
