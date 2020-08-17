package _392IsSubsequence

func isSubsequence(s string, t string) bool {
	charList := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		charList[s[i]]++
	}

	tag := 0

	for i := 0; i < len(t) && tag < len(s); i++ {

		if t[i] == s[tag] {
			if _, ok := charList[t[i]]; ok {
				charList[t[i]]--
			}
			tag++
		}

	}

	for _, v := range charList {
		if v > 0 {
			return false
		}
	}

	return true
}
