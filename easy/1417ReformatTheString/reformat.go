package reformat

func reformat(s string) string {
	nums := make([]byte, 0)
	letters := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			letters = append(letters, s[i])
		} else {
			nums = append(nums, s[i])
		}
	}
	diff := len(nums) - len(letters)
	if diff > 1 || diff < -1 {
		return ""
	}
	return format(nums, letters)
}

func format(nums, letters []byte) string {
	numLen := len(nums)
	letterLen := len(letters)
	result := make([]byte, 0, numLen+letterLen)
	if numLen > letterLen {
		for i := 0; i < numLen; i++ {
			result = append(result, nums[i])
			if i < letterLen {
				result = append(result, letters[i])
			}
		}
	} else {
		for i := 0; i < letterLen; i++ {
			result = append(result, letters[i])
			if i < numLen {
				result = append(result, nums[i])
			}
		}
	}
	return string(result)
}
