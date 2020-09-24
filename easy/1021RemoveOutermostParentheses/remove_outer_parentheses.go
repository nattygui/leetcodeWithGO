package removeouterparentheses

func removeOuterParentheses(S string) string {
	result := ""
	left := 0
	for i := 0; i < len(S); i++ {
		switch S[i] {
		case '(':
			left++
			if left > 1 {
				result += "("
			}
		case ')':
			if left != 1 {
				result += ")"
			}
			left--
		}
	}
	return result
}
