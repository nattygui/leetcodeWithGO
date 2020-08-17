package _035SearchInsertPosition

// 直接遍历法
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// 二分查找法
func searchInsertBinary(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high - low) / 2
		if nums[mid] > target {
			high = mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}