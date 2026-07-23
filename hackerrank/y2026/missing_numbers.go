package y2026

import (
	"fmt"
	"sort"
)





func MissingNumbers(arr []int32, brr []int32) []int32 {
    set := make(map[int32]int32, len(arr))
    for _, v := range arr {
        set[v] ++
    }
    resS := map[int32]struct{}{}
    for _, v := range brr {
        if set[v] == 0 {
            resS[v] = struct{}{}
        } else {
            set[v] --
        }
    }

    fmt.Println(resS)

    keys := make(sortableInt32, 0, len(resS))
    for k := range resS {
        keys = append(keys, k)
    }
    sort.Sort(keys)
    return keys 
}

type sortableInt32 []int32

func (f sortableInt32) Len() int {
	return len(f)
}

func (f sortableInt32) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f sortableInt32) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
