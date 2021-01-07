package temperatures

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	minStack := make([]int, 0)

	for i, v := range T {
		for len(minStack) != 0 {
			index := minStack[len(minStack)-1]
			if T[index] > v {
				break
			}
			res[index] = index - i
			minStack = minStack[0 : len(minStack)-1]
		}
		minStack = append(minStack, i)
	}

	return res
}
