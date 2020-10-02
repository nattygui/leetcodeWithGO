package findocurrences

func findOcurrences(text string, first string, second string) []string {
	result := make([]string, 0)
	i := 0
	j := 0
	lastWord := ""
	findFlag := false
	for i < len(text) {
		if text[i] == ' ' {
			i++
			continue
		}
		j = i
		for j < len(text) && text[j] != ' ' {
			j++
			continue
		}

		if findFlag {
			result = append(result, text[i:j])
			findFlag = false
		}

		if lastWord == first && text[i:j] == second {
			findFlag = true
		}

		lastWord = text[i:j]
		i = j
	}
	return result
}
