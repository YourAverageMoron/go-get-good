package main

import "unicode/utf8"

func ValidAnagram(s string, t string) bool {
    // NOTE: IT WOULD BE BETTER TO USE 2 FOR LOOPS, BUT THIS APPROACH LET ME LEARN MORE ABOUT RUNES IN GO
	if utf8.RuneCountInString(s) != utf8.RuneCountInString(t) {
		return false
	}
	s_rune := []rune(s)
	t_rune := []rune(t)
	m := make(map[rune]int)
	for i := range s {
		char_s := s_rune[i]
		_, ok_s := m[char_s]
		if ok_s {
			m[char_s]++
		} else {
			m[char_s] = 1
		}
		char_t := t_rune[i]
		_, ok_t := m[char_t]
		if ok_t {
			m[char_t]--
		} else {
			m[char_t] = -1
		}
	}

   for _, val := range m {
        if val != 0 {
            return false
        }
   }
   return true
}
