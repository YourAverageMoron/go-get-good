package y2026

func CountTriplets(arr []int64, r int64) int64 {

	m2 := make(map[int64]int64, len(arr))
	m3 := make(map[int64]int64, len(arr))

	count := int64(0)

	for _, v := range arr {
		count += m3[v]
		m3[v*r] += m2[v]
		m2[v*r] += 1
	}

	return count
}
