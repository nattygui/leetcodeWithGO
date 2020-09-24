package heightchecker

func heightChecker(heights []int) int {
	nums := make([]int, 101)
	for _, height := range heights {
		nums[height]++
	}

	result := 0
	heightIndex := 0
	index := 0
	for index <= 101 {
		if nums[index] == 0 {
			index++
			continue
		}
		for i := 0; i < nums[index]; i++ {
			if heights[heightIndex+i] != index {
				result++
			}
		}

		heightIndex += nums[index]
		index++
	}
	return result
}
