package subsets

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	emptySet := make([]int, 0)
	result = append(result, emptySet)
	getAll(nums, emptySet, &result)
	return result
}

func getAll(nums []int, oneSet []int, result *[][]int) {
	if len(nums) == 0 {
		return
	}
	for i, num := range nums {
		tempSet := make([]int, len(oneSet)+1)
		copy(tempSet, oneSet)
		tempSet[len(oneSet)] = num
		*result = append(*result, tempSet)
		getAll(nums[i+1:], tempSet, result)
	}
}
