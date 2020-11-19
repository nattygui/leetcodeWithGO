package buildarray

func buildArray(target []int, n int) []string {
	result := make([]string, 0)
	num := 1
	index := 0
	for num <= n && index < len(target) {
		if num < target[index] {
			result = append(result, "Push")
			result = append(result, "Pop")
		} else {
			result = append(result, "Push")
			index++
		}
		num++
	}
	return result
}
