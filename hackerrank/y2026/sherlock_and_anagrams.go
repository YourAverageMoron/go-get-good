package y2026

import (
	"sort"
	"strings"
)


func sortString(s string) string {
	arr := strings.Split(s, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}


func SherlockAndAnagrams(s string) int32 {
	// Write your code here
	m := map[string]int32{}
	for l := 0; l <= len(s); l++ {
		for i := range s {
			if i+l >= len(s) {
				break
			}
			sub := sortString(s[i : i+l+1])
			m[sub] += 1
		}
	}

	count := int32(0)
	for _, v := range m {
		if v > 1 {
			count += v * (v-1) / 2
		}
	}

	return count
}
