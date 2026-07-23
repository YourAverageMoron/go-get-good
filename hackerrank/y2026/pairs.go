package y2026


func Pairs(k int32, arr []int32) int32 {
	m := make(map[int32]bool, len(arr))

	count := int32(0)
	for _, v := range arr {
		if m[v+k]{
			count++
		}
        if m[v-k] {
            count ++
        }
		m[v] = true
	}
	return count
}
