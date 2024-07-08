package main

func ContainerWithMostWater(height []int) int {
	l := 0
	r := len(height) - 1
	maxV := 0
	for l != r {
		var currH int
		if height[l] > height[r] {
			currH = height[r]
		} else {
			currH = height[l]
		}
		v := currH * (r - l)
		if v > maxV {
			maxV = v
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxV
}

func ContainerWithMostWaterSimple(height []int) int {
	// NOTE: 0(n^2)
	maxV := 0
	for l := range height {
		r := len(height) - 1
		for l != r {
			var currH int
			if height[l] > height[r] {
				currH = height[r]
			} else {
				currH = height[l]
			}
			v := currH * (r - l)
			if v > maxV {
				maxV = v
			}
			r--
		}
	}
	return maxV
}
