package main

import (
	"unicode/utf8"
)

func ValidParentheses(s string) bool {
	stack := make([]rune, utf8.RuneCountInString(s))
	i := -1
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			i++
			stack[i] = c
            continue
		}
		if i == -1 {
            return false
		}
		if c == ')' {
			if stack[i] != '(' {
				return false
			}
		}
		if c == ']' {
			if stack[i] != '[' {
				return false
			}
		}
		if c == '}' {
			if stack[i] != '{' {
				return false
			}
		}
		i--
	}
    if i >= 0{
        return false
    }
	return true
}
