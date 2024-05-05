package quicksort


func Quicksort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivotIndex := partition(arr, low, high)
	quicksort(arr, low, pivotIndex-1) 
	quicksort(arr, pivotIndex+1, high)

}

func partition(arr []int, low, high int) int {
	pivot := arr[high] 
	idx := low - 1

	for i := low; i <= high; i++ {
		if arr[i] < pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	idx++
	arr[high], arr[idx] = arr[idx], pivot
	return idx
}
