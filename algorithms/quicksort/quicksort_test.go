package quicksort

import (
	"fmt"
	"testing"
)

func TestQuicksort(t *testing.T) {
	arr := []int{12, 4, 45, 6, 8, 2, 13}
	fmt.Println("Unsorted array:", arr)
	Quicksort(arr)
	fmt.Println("Sorted array:", arr)
}
