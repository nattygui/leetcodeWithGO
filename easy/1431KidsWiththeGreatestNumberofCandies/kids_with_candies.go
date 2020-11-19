package kidswithcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandi := 0
	for _, num := range candies {
		if maxCandi < num {
			maxCandi = num
		}
	}

	result := make([]bool, 0, len(candies))
	for _, num := range candies {
		if num+extraCandies >= maxCandi {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
