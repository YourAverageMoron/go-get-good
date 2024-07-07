package main


func TopKFrequent(nums []int, k int) []int {
	freqs := make([][]int, len(nums))
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	for num, freq := range m {
		freqs[freq-1] = append(freqs[freq-1], num)
	}

	res := []int{}
	for i := len(freqs) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, freqs[i]...)
	}
	return res
}
