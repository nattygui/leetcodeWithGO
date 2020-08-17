package _290WordPattern

import "strings"

func wordPattern(pattern string, str string) bool {
	wordList := strings.Split(str, " ")
	if len(wordList) != len(pattern) {
		return false
	}

	strByte := make(map[string]byte, 0)
	byteStr := make(map[byte]string, 0)

	for i := 0; i < len(wordList); i++ {
		if v, ok := strByte[wordList[i]]; ok {
			if v != pattern[i] {
				return false
			}
		} else {
			strByte[wordList[i]] = pattern[i]
		}

		if v, ok := byteStr[pattern[i]]; ok {
			if v != wordList[i] {
				return false
			}
		} else {
			byteStr[pattern[i]] = wordList[i]
		}
	}
	return true
}
