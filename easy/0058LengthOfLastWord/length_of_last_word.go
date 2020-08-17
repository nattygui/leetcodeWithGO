package _058LengthOfLastWord

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	index := len(s) - 1
	for index >= 0 && s[index] == ' ' {
		index--
	}

	for index >= 0 && s[index] != ' ' {
		res++
		index--
	}
	return res
}