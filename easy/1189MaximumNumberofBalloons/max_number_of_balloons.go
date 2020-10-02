package maxnumberofballoons

import "math"

func maxNumberOfBalloons(text string) int {
	word := "balloon"
	lettersMap := make(map[byte]int)
	initFindMap := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		lettersMap[word[i]]++
		initFindMap[word[i]] = 0
	}

	for i := 0; i < len(text); i++ {
		if _, ok := initFindMap[text[i]]; ok {
			initFindMap[text[i]]++
		}
	}

	minMultiple := math.MaxInt64
	for key, value := range initFindMap {
		temp := value / lettersMap[key]
		if minMultiple > temp {
			minMultiple = temp
		}
	}
	return minMultiple
}
