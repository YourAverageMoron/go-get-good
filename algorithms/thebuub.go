package main

import "fmt"

func BubbleSort(arr []int) []int {
	count := 0

	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			tmp := arr[i]
			if arr[i] > arr[i+1] {
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
			count++
		}
	}
	fmt.Println(count)
	return arr
}
