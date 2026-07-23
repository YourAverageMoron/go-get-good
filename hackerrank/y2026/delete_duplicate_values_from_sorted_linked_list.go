package y2026

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func DeleteDuplicateValuesFromSortedLinkedList(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
    head := llist
    for llist != nil  && llist.next != nil {
        if llist.data == llist.next.data {
            llist.next = llist.next.next
        } else { 
            llist = llist.next 
        }
    }
    return head
}
