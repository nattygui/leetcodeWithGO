package _389FindTheDifference

func findTheDifference(s string, t string) byte {
	charList := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		charList[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		charList[t[i]]--
	}

	for k, v := range charList {
		if v == -1 {
			return k
		}
	}

	return ' '
}
