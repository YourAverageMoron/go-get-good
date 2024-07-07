package main

func TwoSumIIInputArraySorted(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for {
		val := numbers[start] + numbers[end]
		if val == target {
			return []int{start + 1, end + 1}
		}
		if val > target {
			end--
			continue
		}
		if val < target {
			start++
			continue
		}
	}
}
