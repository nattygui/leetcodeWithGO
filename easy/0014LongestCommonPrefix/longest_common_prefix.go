package _014LongestCommonPrefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLenStr := ""

	// 获取最小的字符串长度
	for index, _ := range strs {
		if index == 0 {
			minLenStr = strs[index]
			continue
		}

		if len(minLenStr) > len(strs[index]) {
			minLenStr = strs[index]
		}
	}
	if minLenStr == "" {
		return ""
	}

	for minLenStr != "" {
		satisfied := 0
		for _, v := range strs {
			if strings.HasPrefix(v, minLenStr) {
				satisfied++
				continue
			} else {
				minLenStr = minLenStr[:len(minLenStr)-1]
				break
			}
		}
		if satisfied == len(strs) {
			return minLenStr
		}
	}
	return ""
}