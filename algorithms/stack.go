package main

type StackNode[T any] struct {
	value T
	next  *StackNode[T]
}

type Stack[T any] struct {
	head *StackNode[T]
}


func (s *Stack[T]) Push(value T){
    newNode := &StackNode[T]{value, s.head}
    s.head = newNode
}

func (s *Stack[T]) Pop() T {
    if s.head == nil{
        panic("Stack empty")
    }
    h := s.head
    s.head = s.head.next
    return h.value
}

func (s *Stack[T]) Peek() T {
    if s.head == nil{
        panic("Stack empty")
    }
    return s.head.value
}
