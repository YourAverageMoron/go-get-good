package y2026

func FreqQuery(queries [][]int32) []int32 {
	m := make(map[int32]int32, len(queries))
	freqs := make(map[int32]int32)

	res := []int32{}

	for _, q := range queries {
		val := q[1]
		if q[0] == 1 {
			if freqs[m[val]] > 0 {
				freqs[m[val]]--
			}
			m[val]++
			freqs[m[val]]++
		}
		if q[0] == 2 {
			if m[val] > 0 {
				if freqs[m[val]] > 0 {
					freqs[m[val]]--
				}
				m[val]--
				freqs[m[val]]++

			}
		}
		if q[0] == 3 {
            if freqs[val] > 0 {
                res = append(res, 1)
            } else {
                res = append(res, 0)
            }

		}
	}
	return res
}
