package smallernumbersthancurrent

func smallerNumbersThanCurrent(nums []int) []int {
	// 统计所有的元素，个数
	static := make([]int, 101)
	for _, num := range nums {
		static[num]++
	}

	for i := 1; i < len(static); i++ {
		static[i] += static[i-1]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = static[nums[i]]
	}
	return nums
}
