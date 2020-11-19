package maxlength

func maxLengthBetweenEqualCharacters(s string) int {
	maxLength := -1
	charsMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if v, ok := charsMap[s[i]]; !ok {
			charsMap[s[i]] = i
		} else {
			if maxLength < i-v-1 {
				maxLength = i - v - 1
			}
		}
	}
	return maxLength
}
