package main

func GenerateParentheses(n int) []string {
	res := &[]string{}
	generateParentheses(n, 0, 0, []rune{}, res)
	return *res

}

func generateParentheses(n int, open int, closed int, s []rune, res *[]string) {
	if open == n && closed == n {
		*res = append(*res, string(s))
	}

	if open > closed {
		newClosed := closed + 1
		s1 := append(s, ')')
		generateParentheses(n, open, newClosed, s1, res)
	}

	if open <= n {
		open++
		s2 := append(s, '(')
		generateParentheses(n, open, closed, s2, res)
	}
}
