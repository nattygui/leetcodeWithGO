package dominant

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxInt, seMaxInt, maxIndex := getSecondValue(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		v := nums[i]
		if maxInt < v {
			if maxInt > seMaxInt {
				maxInt, seMaxInt = v, maxInt
			} else {
				maxInt = v
			}
			maxIndex = i
		} else if seMaxInt < v && v != maxInt {
			seMaxInt = v
		}
	}
	if maxInt >= seMaxInt*2 {
		return maxIndex
	}
	return -1
}

func getSecondValue(value1, value2 int) (int, int, int) {
	if value1 > value2 {
		return value1, value2, 0
	}
	return value2, value1, 1
}
