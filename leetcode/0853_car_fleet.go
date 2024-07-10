package main

import (
	"sort"
)

func CarFleet(target int, position []int, speed []int) int {
	cars := make([][]int, len(position))
	for i := range position {
		cars[i] = []int{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})
	s := make([]float64, len(position))
	l := 0
	for _, car := range cars {
		time_to_target := float64(target - car[0]) / float64(car[1])
		if l == 0 || time_to_target > s[l-1] {
			s[l] = time_to_target
			l++
		}
	}
	return l
}
