package y2026

import "fmt"

func IceCreamParlour(cost []int32, money int32) {
	m := make(map[int32]int, len(cost))
	for i, c := range cost {
		if j, ok := m[money-c]; ok {
			fmt.Println(j+1, i+1)
			return
		}
		m[c] = i
	}
}
