package quicksort

import "sync"

type SafeArr struct {
	mu  sync.Mutex
	arr []int
}

func QuicksortParallel(arr []int) {
	sarr := &SafeArr{arr: arr}
	var wg sync.WaitGroup
	quicksortParallel(sarr, 0, len(arr)-1, &wg)
}

func quicksortParallel(arr *SafeArr, low, high int, wg *sync.WaitGroup) {
	if low >= high {
		return
	}
	pivotIndex := partitionParallel(arr, low, high)
	wg.Add(2)
	go func() {
		defer wg.Done()
		quicksortParallel(arr, low, pivotIndex-1, wg)
	}()
	go func() {
		defer wg.Done()
		quicksortParallel(arr, pivotIndex+1, high, wg)
	}()

}

func partitionParallel(sa *SafeArr, low, high int) int {
	pivot := sa.arr[high]
	idx := low - 1
	for i := low; i <= high; i++ {
		sa.mu.Lock()
		if sa.arr[i] < pivot {
			idx++
			sa.arr[i], sa.arr[idx] = sa.arr[idx], sa.arr[i]
		}
		sa.mu.Unlock()
	}
	idx++
	sa.mu.Lock()
	sa.arr[high], sa.arr[idx] = sa.arr[idx], pivot
	sa.mu.Unlock()
	return idx
}

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
