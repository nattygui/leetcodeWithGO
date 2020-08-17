package _026RemoveDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	hasDelete := 0
	var s = make(map[int]int, 0)
	for i, v := range temp {
		if _, ok := s[v]; ok {
			s[v]++
			nums = append(nums[:i - hasDelete], nums[i - hasDelete + 1:]...)
			hasDelete++
		} else {
			s[v] = 0
			s[v]++
		}
	}

	return len(nums)
}

func removeDuplicatesFastsort(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] < nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j+1
}
