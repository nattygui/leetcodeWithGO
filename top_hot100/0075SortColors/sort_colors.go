package sortcolors

func sortColors(nums []int) {
	count0, count1 := 0, 0
	for i, num := range nums {
		if num == 0 {
			nums[i], nums[count0] = nums[count0], nums[i]
			if count0 < count1 {
				nums[i], nums[count1] = nums[count1], nums[i]
			}
			count0++
			count1++
		} else if num == 1 {
			nums[i], nums[count1] = nums[count1], nums[i]
			count1++
		}
	}
}
