package removeduplicates

func removeDuplicates(S string) string {
	str := []byte(S)

	startIndex := 0
	endIndex := 0
	for {
		for startIndex+1 < len(str) && str[startIndex] != str[startIndex+1] {
			startIndex++
		}
		if startIndex >= len(str)-1 {
			break
		}
		endIndex = startIndex + 2
		str = append(str[0:startIndex], str[endIndex:]...)
		startIndex = 0
		endIndex = 0
	}
	return string(str)
}
