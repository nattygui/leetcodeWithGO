package _189RotateArray

func rotate(nums []int, k int)  {
	if len(nums) <= 1 || k <= 0 || k >= len(nums) {
		return
	}
	exchange(nums, 0, len(nums)-k, len(nums))
}

func exchange(nums []int, start int, mid int, end int) {
	if (mid - start) == (end - mid) {
		for i := 0; i < (mid-start); i++ {
			temp := 0
			temp = nums[start+i]
			nums[start+i] = nums[mid+i]
			nums[mid+i] = temp
		}
		return
	} else if (mid - start) < (end - mid) {
		for i := 0; i < (mid-start); i++ {
			temp := 0
			temp = nums[start+i]
			nums[start+i] = nums[mid+i]
			nums[mid+i] = temp
		}
		exchange(nums, mid, mid+mid-start, end)
	} else {
		for i := 0; i < (end - mid); i++ {
			temp := 0
			temp = nums[start+i]
			nums[start+i] = nums[mid+i]
			nums[mid+i] = temp
		}
		exchange(nums, start+end-mid, mid, end)
	}
}
