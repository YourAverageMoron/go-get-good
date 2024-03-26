package main

import (
	"fmt"
	"reflect"
)

type SingleNode[T any] struct {
	value T
	next  *SingleNode[T]
}

type SinglyLinkedList[T any] struct {
	head *SingleNode[T]
}

func (l *SinglyLinkedList[T]) Append(value T) {
	newNode := &SingleNode[T]{value, nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (l *SinglyLinkedList[T]) Prepend(value T) {
	newNode := &SingleNode[T]{value, l.head}
	l.head = newNode
}

func (l SinglyLinkedList[T]) Delete(index int) {
	curr := l.head
	for i := 0; i <= index; i++ {
		if curr == nil {
			return
		}
		curr = curr.next
	}
	curr.next = curr.next.next
}

func (l *SinglyLinkedList[T]) Display() {
	curr := l.head
	if curr == nil {
		return
	}
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

type DoublyNode[T any] struct {
	value      T
	prev, next *DoublyNode[T]
}

type DoublyLinkedList[T any] struct {
	head, tail *DoublyNode[T]
}

func (l *DoublyLinkedList[T]) Append(value T) {
	newNode := &DoublyNode[T]{value, l.tail, nil}
	if l.head == nil {
		l.head = newNode
	}
	if l.tail != nil {
		l.tail.next = newNode
	}
	l.tail = newNode
}

func (l *DoublyLinkedList[T]) Prepend(value T) {
	newNode := &DoublyNode[T]{value, nil, l.head}
	if l.head != nil {
		l.head.prev = newNode
	}
	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
}

func (l DoublyLinkedList[T]) Length() int {
	curr := l.head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (l *DoublyLinkedList[T]) Get(index int) T {
	curr := l.head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.next
	}
	if curr == nil {
		panic("Index out of range")
	}
	return curr.value
}

func (l *DoublyLinkedList[T]) InsertAt(value T, index int) {
	curr := l.head
	for i := 1; i < index && curr != nil; i++ {
		curr = curr.next
	}
	if curr == nil {
		panic("Index out of range")
	}
	newNode := &DoublyNode[T]{value, curr, curr.next}
	curr.next.prev = newNode
	curr.next = newNode
}

func (l *DoublyLinkedList[T]) RemoveAt(index int) {
	curr := l.head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.next
	}
	if curr == nil {
		panic("Index out of range")
	}
	prev := curr.prev
	next := curr.next
	if curr == l.head {
		l.head = next
		next.prev = nil
		return
	}
	if curr == l.tail {
		l.tail = prev
		prev.next = nil
		return
	}

	prev.next = curr.next
	next.prev = curr.prev
}

func (l DoublyLinkedList[T]) Remove(value T) {
	curr := l.head
	for curr != nil {
		if reflect.ValueOf(curr.value).Interface() == reflect.ValueOf(value).Interface() {
			prev := curr.prev
			next := curr.next
			if curr == l.head {
				l.head = next
				next.prev = nil
				return
			}
			if curr == l.tail {
				l.tail = prev
				prev.next = nil
				return
			}
			prev.next = curr.next
			next.prev = curr.prev
		}
		curr = curr.next
	}
}

func (l *DoublyLinkedList[T]) Display() {
	curr := l.head
	if curr == nil {
		return
	}
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}
