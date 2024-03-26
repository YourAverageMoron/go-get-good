package main

func main() {
	//CRYSTAL
	// fmt.Println("Hello, World!")
	// breaks := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	// fmt.Println(CystalBall(breaks))

	//BUBBLE
	// arr := []int{1, 3, 7, 4, 2}
	// fmt.Println(BubbleSort(arr))

	l := DoublyLinkedList[int]{nil, nil}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(3)
	l.Remove(3)

	l.Display()
}
