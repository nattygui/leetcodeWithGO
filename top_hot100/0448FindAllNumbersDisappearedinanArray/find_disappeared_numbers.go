package finddisappearednumbers

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	i := 0
	for i < len(nums) {
		if nums[i] == i+1 || nums[nums[i]-1] == nums[i] {
			i++
			continue
		}
		nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
	}

	for i, v := range nums {
		if i+1 != v {
			result = append(result, i+1)
		}
	}
	return result
}
