package main

import (
	"fmt"
	"math"
)

func CystalBall(breaks []bool) int {
	interval := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := 0
	for ; i < len(breaks); i += interval {
		fmt.Println(i)
		if breaks[i] {
			break
		}
	}
	i -= interval
	for j := 0; j <= interval && i+j < len(breaks); j++ {
		if breaks[i+j] {
			return i + j
		}
	}
	fmt.Println(i)
	return 0
}
