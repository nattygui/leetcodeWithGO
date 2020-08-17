package findpivotindex

func pivotIndex(nums []int) int {
	nums = append([]int{0}, nums...)
	nums = append(nums, 0)
	index := 1
	leftSum := addALL(nums[0:index])
	rightSum := addALL(nums[index+1:])
	for index < len(nums)-1 {
		if leftSum == rightSum {
			return index - 1
		}
		leftSum += nums[index]
		rightSum -= nums[index+1]
		index++
	}
	return -1
}

func addALL(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
