package _704BinarySearch

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	mid := 0
	temp := 0
	for i <= j {
		mid = (i + j) / 2
		temp = nums[mid]
		if temp == target {
			return mid
		} else if temp < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}
