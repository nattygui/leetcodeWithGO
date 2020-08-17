package _520DetectCapital

func detectCapitalUse(word string) bool {
	count := 0
	flag := false

	for i, v := range word {
		if v >= 'A' && v <= 'Z' {
			count++
			if i == 0 {
				flag = true
			}
		}
	}

	if count == 0 || count == len(word) {
		return true
	}
	if count == 1 && flag == true {
		return true
	}
	return false

}