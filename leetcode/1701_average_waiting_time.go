package main

func AverageWaitingTime(customers [][]int) float64 {
	time := 0
	total := 0
	totalC := 0
	for _, customer := range customers {
		if time < customer[0] {
			time = customer[0]
		}
        time += customer[1]
        total += time - customer[0]
		totalC++
	}
	return float64(total) / float64(totalC)
}
