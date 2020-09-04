package shortesttochar

func shortestToChar(S string, C byte) []int {
	findCIndex := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			findCIndex = append(findCIndex, i)
		}
	}

	res := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			res[i] = 0
			continue
		}
		res[i] = findShortest(i, findCIndex)
	}
	return res
}

func findShortest(index int, cIndex []int) int {
	for i := 0; i < len(cIndex); i++ {
		if cIndex[i] > index {
			if i == 0 {
				return cIndex[i] - index
			}
			return mini(cIndex[i]-index, index-cIndex[i-1])
		}

		if i == len(cIndex)-1 {
			return index - cIndex[i]
		}

		if cIndex[i+1] < index {
			continue
		}

		return mini(cIndex[i+1]-index, index-cIndex[i])

	}
	return 0
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}
