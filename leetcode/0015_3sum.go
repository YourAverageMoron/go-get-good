package main

import "sort"

func ThreeSumTwoPointer(nums []int) [][]int {
    results := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue //To prevent the repeat
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
    return results
}

func ThreeSumMap(nums []int) [][]int {
	freqs := make(map[int]int)
	res := [][]int{}
	seen := make(map[int]bool)

	for _, n := range nums {
		freqs[n]++
	}
	for n1 := range freqs {
		if s := seen[n1]; s {
			continue
		}

		tempSeen := make(map[int]bool)
		for n2 := range freqs {
			if s := seen[n2]; s {
				continue
			}
			if s := tempSeen[n2]; s {
				continue
			}
			if n1 == n2 && freqs[n1] < 2 {
				continue
			}
			target := -(n1 + n2)
			if s := seen[target]; s {
				continue
			}
			if s := tempSeen[target]; s {
				continue
			}
			count := 1
			if target == n2 {
				count++
			}
			if target == n1 {
				count++
			}
			if x := freqs[target]; x >= count {
				res = append(res, []int{n1, n2, target})
			}
			tempSeen[target] = true
		}
		seen[n1] = true
	}
	return res
}
