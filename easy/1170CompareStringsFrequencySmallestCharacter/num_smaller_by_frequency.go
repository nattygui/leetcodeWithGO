package numsmallerbyfrequency

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	wordsNum := make([]int, len(words))

	for i := 0; i < len(words); i++ {
		wordsNum[i] = getMinLetterTimes(words[i])
	}

	sort.Ints(wordsNum)
	result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		num := getMinLetterTimes(queries[i])
		result[i] = getNum(wordsNum, num)
	}
	return result
}

func getMinLetterTimes(word string) int {
	letterMap := make(map[byte]int)

	c := byte('z')
	for i := 0; i < len(word); i++ {
		if c > word[i] {
			c = word[i]
		}
		letterMap[word[i]]++
	}
	return letterMap[c]
}

func getNum(wordsNum []int, num int) int {
	i := 0
	j := len(wordsNum) - 1

	if num < wordsNum[i] {
		return len(wordsNum)
	}

	if num >= wordsNum[len(wordsNum)-1] {
		return 0
	}

	for i <= j {
		mid := (i + j) / 2
		if num >= wordsNum[mid] && num < wordsNum[mid+1] {
			return len(wordsNum) - mid - 1
		}
		if num >= wordsNum[mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return 0
}
