package numequivdominopairs

import "strconv"

func numEquivDominoPairs(dominoes [][]int) int {
	namesMap := make(map[string]int)
	for _, domino := range dominoes {
		namesMap[getDominoName(domino)]++
	}

	result := 0
	for _, v := range namesMap {
		result += v * (v - 1) / 2
	}

	return result
}

func getDominoName(domino []int) string {
	if domino[0] < domino[1] {
		return strconv.Itoa(domino[0]) + "|" + strconv.Itoa(domino[1])
	}
	return strconv.Itoa(domino[1]) + "|" + strconv.Itoa(domino[0])
}
