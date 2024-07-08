package main


func TrappingRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l := 0
	r := len(height) - 1
	lMax := height[l]
	rMax := height[r]
	res := 0

	for l < r {
		if lMax < rMax {
            l ++
            if height[l] > lMax  {
                lMax = height[l]
            }
            res += lMax - height[l]
		} else {
            r --
            if height[r] > rMax {
                rMax = height[r]
            }
            res += rMax - height[r]
        }
	}
    return res
}

func TrappingRainWaterInefficient(height []int) int {
	total := 0
	currV := 0

	l := 0
	currH := height[l]
	r := 0
	for {
		if l == len(height)-1 {
			return total
		}
		r++
		if height[r] >= currH {
			total += currV
			currV = 0
			l = r
			currH = height[l]
			continue
		}
		if r == len(height)-1 {
			currV = 0
			r = l
			currH--
			continue
		}
		currV += currH - height[r]
	}
}
