package distributecandies

func distributeCandies(candies int, numPeople int) []int {
	result := make([]int, numPeople)
	index := 0
	times := 0
	for candies > 0 {
		if index == numPeople {
			index = 0
			times++
		}

		if candies >= times*numPeople+index+1 {
			result[index] += times*numPeople + index + 1
			candies -= times*numPeople + index + 1
			index++
		} else {
			result[index] += candies
			candies = 0
		}
	}
	return result
}
