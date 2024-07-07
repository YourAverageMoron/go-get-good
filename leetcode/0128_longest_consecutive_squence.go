package main

import "fmt"

func LongestConsecutiveSequence(nums []int) int {
	m := make(map[int]int)
	maxL := 0
	for _, n := range nums {
		if _, exists := m[n]; exists {
			continue
		}
		prevL := m[n-1]
		nextL := m[n+1]
		l := 1 + nextL + prevL
		if l > maxL {
			maxL = l
			fmt.Println(n)
		}
		m[n] = l
		m[n-prevL] = l
		m[n+nextL] = l
	}
	return maxL
}
