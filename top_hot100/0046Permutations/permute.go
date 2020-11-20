package permute

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	walk(nums, nil, &result, len(nums))
	return result
}

func walk(nums []int, oneResult []int, result *[][]int, length int) {
	if oneResult == nil {
		oneResult = make([]int, 0, length)
	}
	for i, num := range nums {
		tempResult := append(oneResult, num)
		nextNums := make([]int, 0, len(nums)-1)
		nextNums = append(nextNums, nums[:i]...)
		nextNums = append(nextNums, nums[i+1:]...)
		if len(tempResult) == length {
			// 添加一个结果的时候，由于slice 的特性，需要重新创建一个切片
			temp := make([]int, length)
			copy(temp, tempResult)
			*result = append(*result, temp)
		} else {
			walk(nextNums, tempResult, result, length)
		}
	}
}
