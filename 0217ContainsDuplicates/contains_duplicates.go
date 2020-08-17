package _217ContainsDuplicates

func containsDuplicate(nums []int) bool {
	isExists := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := isExists[nums[i]]; ok {
			return true
		} else {
			isExists[nums[i]] = 0
		}
	}
	return false
}
