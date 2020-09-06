package reorder

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	res := make([]string, len(logs))
	indexRes := len(logs) - 1
	wordstr := make([]string, 0)

	for i := len(logs) - 1; i >= 0; i-- {
		if isDigitContents(logs[i]) {
			res[indexRes] = logs[i]
			indexRes--
		} else {
			wordstr = append(wordstr, logs[i])
		}
	}

	sort.Sort(logsFile(wordstr))
	copy(res, wordstr)
	return res
}

func isDigitContents(str1 string) bool {
	strs := strings.Split(str1, " ")
	strs = strs[1:]

	for _, str := range strs {
		for _, s := range str {
			if s >= '0' && s <= '9' {
				continue
			}
			return false
		}
	}
	return true
}

type logsFile []string

func (lf logsFile) Len() int {
	return len(lf)
}

func (lf logsFile) Swap(i, j int) {
	lf[i], lf[j] = lf[j], lf[i]
}

func (lf logsFile) Less(i, j int) bool {
	str1, str2 := lf[i], lf[j]

	index1 := strings.IndexByte(str1, ' ')
	index2 := strings.IndexByte(str2, ' ')

	Identifier1 := str1[:index1]
	Identifier2 := str2[:index2]
	content1 := str1[index1:]
	content2 := str2[index2:]

	if content1 != content2 {
		return stringLess(content1, content2)
	}

	return stringLess(Identifier1, Identifier2)
}

func stringLess(str1, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)

	lens := min(len1, len2)
	for i := 0; i < lens; i++ {
		if str1[i] < str2[i] {
			return true
		}
		if str1[i] > str2[i] {
			return false
		}
	}
	if lens < len1 {
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
