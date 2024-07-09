package main

import (
	"strconv"
)

func ReversePolishNotation(tokens []string) int {
	s := make([]int, len(tokens))
	l := 0
	for _, c := range tokens {
		if c == "+" {
			a := s[l-1]
			l--
			s[l-1] = a + s[l-1]
		} else if c == "-" {
			a := s[l-1]
			l--
			s[l-1] = s[l-1] - a

		} else if c == "*" {
			a := s[l-1]
			l--
			s[l-1] = s[l-1] * a
		} else if c == "/" {

			a := s[l-1]
			l--
			s[l-1] = s[l-1] / a
		} else {
			s[l], _ = strconv.Atoi(c)
			l++
			continue
		}
	}

	return s[l-1]
}
