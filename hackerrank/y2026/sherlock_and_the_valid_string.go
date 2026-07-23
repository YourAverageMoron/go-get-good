package y2026

import (
	"math"
)

func SherlockAndTheValidString(s string) string {
	m := make(map[rune]int32, len(s))
	for _, c := range s {
		m[c]++
	}

    freqs := map[int32]int32{}
    for _, f := range m {
        freqs[f] ++
    }

	if len(freqs) < 2 {
		return "YES"
	}

	if len(freqs) > 2 {
		return "NO"
	}

	keys := make([]rune, 0, len(m))
	for k := range freqs {
		keys = append(keys, k)
	}

    if (keys[0] == 1 && freqs[keys[0]] == 1) || (keys[1] == 1 && freqs[keys[1]] == 1) {
        return "YES"
    }

    if math.Abs(float64(keys[0] - keys[1])) > 1 || ( freqs[keys[0]] > 1 && freqs[keys[1]] > 1) {
		return "NO"
	}
	return "YES"
}
