package main


type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{value, nil}
	if q.tail == nil {
		q.tail = newNode
		q.head = newNode
        return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (l *Queue[T]) Deque() T {
	if l.head == nil {
		panic("Queue is empty")
	}
	h := l.head
	l.head = h.next
	h.next = nil
	return h.value
}

func (l *Queue[T]) Peek() T {
	if l.head == nil {
		panic("Queue is empty")
	}
	return l.head.value
}
