package main

import (
	"fmt"
	"math"
)

type RingBuffer[T any] struct {
	arr    []T
	head   int
	tail   int
	length int
}

func newRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		make([]T, capacity),
		0,
		0,
		0,
	}
}

func (rb *RingBuffer[T]) increaseBufferCapacity() {
	newArr := make([]T, len(rb.arr)*2)
    var index int
	for i := 0; i < rb.length; i++ {
        index = getRelativeIndex(i + rb.head, len(rb.arr))
		fmt.Println(index)
		newArr[i] = rb.arr[index]
	}
    rb.head = 0
    rb.tail = rb.length
	rb.arr = newArr
}

func getRelativeIndex(index int, capacity int) int {
	if index < 0 {
		relativeIndex := capacity - int(math.Abs(float64((index)%6)))
		if relativeIndex == capacity {
			relativeIndex = 0
		}
		return relativeIndex
	}

	return index % capacity
}

func (rb *RingBuffer[T]) Push(item T) {
	if rb.length == len(rb.arr) {
		rb.increaseBufferCapacity()
        fmt.Println(rb)
	}
	index := getRelativeIndex(rb.tail, len(rb.arr))
	rb.arr[index] = item
	rb.tail++
	rb.length++

}

func (rb *RingBuffer[T]) Equeue(item T) {
	if rb.length == len(rb.arr) {
		panic("bbb")
	}
	index := getRelativeIndex(rb.head-1, len(rb.arr))
	rb.arr[index] = item
	rb.head--
	rb.length++
}
