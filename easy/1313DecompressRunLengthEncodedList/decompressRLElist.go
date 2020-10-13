package decompressrlelist

func decompressRLElist(nums []int) []int {
	length := 0
	for i := 0; i < len(nums); i = i + 2 {
		length += nums[i]
	}

	result := make([]int, length)
	index := 0
	for i := 0; i < len(nums); i = i + 2 {
		value := nums[i+1]
		for j := 0; j < nums[i]; j++ {
			result[index] = value
			index++
		}
	}

	return result
}
