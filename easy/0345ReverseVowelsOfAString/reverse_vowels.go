package _345ReverseVowelsOfAString

func reverseVowels(s string) string {
	letterList := []byte(s)

	var temp byte
	i := 0
	j := len(letterList) - 1

	for i < j {
		if isVowel(letterList[i]) && isVowel(letterList[j]) {
			temp = letterList[i]
			letterList[i] = letterList[j]
			letterList[j] = temp
			i++
			j--
		} else if !(isVowel(letterList[i])) {
			i++
		} else {
			j--
		}
	}
	return string(letterList)
}

func isVowel(letter byte) bool {
	switch letter {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	case 'A':
		return true
	case 'E':
		return true
	case 'I':
		return true
	case 'O':
		return true
	case 'U':
		return true
	}
	return false
}