package match

func diStringMatch(S string) []int {
	res := make([]int, len(S))

	inIndex := 0
	deIndex := len(S)

	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			res[i] = inIndex
			inIndex++
		} else if S[i] == 'D' {
			res[i] = deIndex
			deIndex--
		}
	}
	res = append(res, inIndex)
	return res
}
