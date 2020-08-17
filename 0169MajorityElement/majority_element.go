package _169MajorityElement

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 0
	maxNum := nums[0]
	statistical := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := statistical[nums[i]]; ok {
			statistical[nums[i]]++
			if max < statistical[nums[i]] {
				max = statistical[nums[i]]
				maxNum = nums[i]
			}
		} else {
			statistical[nums[i]] = 1
		}
	}
	return maxNum
}

func majorityElementOne(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maj := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if maj == nums[i] {
			count++
		} else if count == 1 {
			maj = nums[i]
		} else {
			count--
		}
	}
	return maj
}