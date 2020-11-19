package generateparenthesis

import "strings"

func generateParenthesis(n int) []string {
	var result []string
	generateAll("", n, n, &result)
	return result
}

func generateAll(curStr string, left int, right int, result *[]string) {
	// 若左括号和右括号全部拼接完成，则追加到结果中
	if left == 0 && right == 0 {
		*result = append(*result, curStr)
		return
	}
	// 若左括号全部拼接，则剩下的全是右括号
	if left == 0 {
		curStr += strings.Repeat(")", right)
		*result = append(*result, curStr)
		return
	}
	// 若左括号剩下的未拼接的小于右括号未拼接的，则有俩种可能，拼接左以及右括号
	if left < right {
		generateAll(curStr+"(", left-1, right, result)
		generateAll(curStr+")", left, right-1, result)
	} else if left == right { // 若左括号剩下的未拼接的等于右括号未拼接的， 则只能进行左括号
		generateAll(curStr+"(", left-1, right, result)
	}
}
