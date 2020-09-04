package uncommonsentences

import "strings"

func uncommonFromSentences(A string, B string) []string {
	wordsMap := make(map[string]int)

	aS := strings.Split(A, " ")
	bS := strings.Split(B, " ")

	for _, a := range aS {
		wordsMap[a]++
	}

	for _, b := range bS {
		wordsMap[b]++
	}

	res := []string{}
	for k, v := range wordsMap {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}
