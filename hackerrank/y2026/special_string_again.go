package y2026

func SpecialStringAgain(n int32, s string) int64 {
	count := int64(0)
	runes := []rune(s)
	for i, c := range runes {
		count++

		j := 1
		for i+j < len(s) && runes[i+j] == c {
			count++
			j++
		}

		if i == 0 || i == len(s)-1 || runes[i-1] == c {
			continue
		}
		l := 1
		r := 1

		selected := runes[i-1]
		for i-l >= 0 && i+r < len(s) && runes[i-l] == selected && runes[i+r] == selected {
			count++
			l++
			r++
		}
	}

	return count
}

// This was super ineffiecient
// func SpecialStringAgain(n int32, s string) int64 {
//
//     count := int64(len(s))
//     for l := 2; l <= len(s); l ++ {
//         for i := range s {
//             if l + i > len(s) {
//                 break
//             }
//             substring := s[i: i + l]
//             if isPalindrome(substring){
//                 count ++
//             }
//         }
//     }
//     return count
// }
//
// func isPalindrome(s string) bool {
//     if len(s) == 0 {
//         return false
//     }
// 	r := len(s) - 1
//     selected := s[0]
// 	for l := range s {
// 		if l >= r {
// 			return true
// 		}
// 		if s[l] != selected || s[r] != selected {
// 			return false
// 		}
// 		r--
// 	}
// 	return true
// }
