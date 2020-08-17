package _383RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	letterList := make([]int, 26)

	for _, v := range ransomNote {
		letterList[v - 'a']--
	}

	for _, v := range magazine {
		letterList[v - 'a']++
	}

	for _, v := range letterList {
		if v < 0 {
			return false
		}
	}
	return true
}