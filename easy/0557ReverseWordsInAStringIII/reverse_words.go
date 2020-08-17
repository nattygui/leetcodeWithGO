package _557ReverseWordsInAStringIII

func reverseWords(s string) string {
	res := []byte(s)

	i, j := 0, 0

	for j < len(s) {
		for j < len(s) && res[j] == ' ' {
			j++
		}
		i = j
		for j < len(s) && res[j] != ' '  {
			j++
		}
		for t := 0; t < (j-i) / 2; t++ {
			res[i+t], res[j-1-t] = res[j-1-t], res[i+t]
		}
	}
	return string(res)
}
