package maxarea

func maxArea(height []int) int {
	i, j, result := 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			result = max(result, height[i]*(j-i))
			i++
		} else {
			result = max(result, height[j]*(j-i))
			j--
		}
	}
	return result
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
