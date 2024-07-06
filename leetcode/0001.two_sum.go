package main

func TwoSumLists(nums []int, target int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSumMap(nums []int, target int) []int {
	m := make(map[int]int)
	for index, n := range nums {
        hashIndex, ok := m[target - n] 
        if ok {
            return []int{index, hashIndex}
        }
        m[n] = index
	}
    return []int{}
}
