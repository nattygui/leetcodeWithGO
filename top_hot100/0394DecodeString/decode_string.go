package decode

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	result := ""
	i := 0
	for {
		if i >= len(s) {
			break
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			result += s[i : i+1]
			i++
		} else {
			endIndex, temp := next(s, i)
			result += temp
			i = endIndex + 1
		}
	}
	return result
}

func next(s string, j int) (int, string) {
	// 获取当前数字
	str := ""
	num := getDigits(s, j)
	if s[j] == '[' {
		j++
	}
	for {
		if j >= len(s) {
			break
		}
		if s[j] == ']' {
			break
		}
		if s[j] >= '0' && s[j] <= '9' {
			endIndex, nextStr := next(s, j)
			j = endIndex + 1
			str += nextStr
		} else {
			str += s[j : j+1]
			j++
		}
	}
	return j, strings.Repeat(str, num)
}

func getDigits(s string, i int) int {
	j := i
	for s[j] >= '0' && s[j] <= '9' {
		j++
	}
	num, _ := strconv.Atoi(s[i:j])
	return num
}
