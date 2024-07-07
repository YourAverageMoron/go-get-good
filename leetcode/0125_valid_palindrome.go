package main

import (
	"unicode/utf8"
)

func ValidPalindrome(s string) bool {
    // NOTE: THIS WOULD PROBS BE BETTER WITH unicode.IsLetter/unicode.IsNumber/unicode.ToLower, but whats the fun in that
	start := 0
	end := utf8.RuneCountInString(s) - 1

	for {
		if start >= end {
			return true
		}
		r1 := rune(s[start])
		if r1 >= 65 && r1 <= 90 {
			r1 += 32
		}
		if (r1 < 97 || r1 > 122) && (r1 < 48 || r1 > 57) {
			start++
			continue
		}
		r2 := rune(s[end])
		if r2 >= 65 && r2 <= 90 {
			r2 += 32
		}
		if (r2 < 97 || r2 > 122) && (r2 < 48 || r2 > 57) {
			end--
			continue
		}
		if r1 != r2 {
			return false
		}
		start++
		end--
	}
}
