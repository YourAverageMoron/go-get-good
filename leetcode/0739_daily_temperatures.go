package main


func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	s := make([][]int, len(temperatures))
	l := 0
	for i := len(temperatures) - 1; i >= 0; i-- {
		for l != 0 && temperatures[i] >= s[l-1][0] {
			l--
		}
		if l == 0 {
			res[i] = 0
		} else {
			res[i] = s[l-1][1] - i
		}
		s[l] = []int{temperatures[i], i}
		l++
	}
	return res
}

