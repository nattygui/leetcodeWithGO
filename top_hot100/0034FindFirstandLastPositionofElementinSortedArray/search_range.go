package searchrange

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	index := -1
	find := false
	for i, j := 0, len(nums)-1; i <= j; {
		mid := (i + j) / 2
		if nums[mid] == target {
			find = true
			index = mid
			break
		}
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	if find {
		i := index
		for i >= 0 {
			if nums[i] == nums[index] {
				i--
			} else {
				result[0] = i + 1
				break
			}
			if i == 0 {
				result[0] = 0
			}
		}
		j := index
		for j <= len(nums)-1 {
			if nums[j] == nums[index] {
				j++
			} else {
				result[1] = j - 1
				break
			}
			if j == len(nums)-1 {
				result[1] = len(nums) - 1
			}
		}
	}

	return result
}
