package main

func WinnerCircularGame(n int, k int) int {
	winner := 0
	for i := 1; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1
}

func WinnerCircularGameOld(n int, k int) int {
	arr := make([]int, n)
	removed := 0
	i := 0
	for {
		steps := 0
		for {
			if i == n {
				i = 0
			}
			if arr[i] == 1 {
				i++
				continue
			}
			if steps+1 == k {
				break
			}
			i++
			steps++
		}
		arr[i] = 1
		if removed == n-1 {
			return i + 1
		}
		removed++
		i++

	}
}
