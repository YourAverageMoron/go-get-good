package main

func GroupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		l := [26]int{}
		for _, c := range str {
			l[c-'a']++
		}
		arr, _ := m[l]
		m[l] = append(arr, str)
	}
	ans := [][]string{}
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}

