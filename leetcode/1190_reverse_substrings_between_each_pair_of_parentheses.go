package main


func ReverseSubstringsBetweenEachPairOfParentheses(s string) string {
	stack := make([]string, (len(s)/3)+2)
	l := 0
	innerS := ""
	for _, c := range s {
		if c == '(' {
			stack[l] = innerS
			l++
			innerS = ""
		} else if c == ')' {
            innerS = ReverseString(innerS)
			l--
			innerS = stack[l] + innerS
		} else {
			innerS += string(c)
		}
	}
	return innerS
}


func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
