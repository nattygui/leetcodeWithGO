package TwoSum

func TwoSum(numbers []int, target int) []int {
	out := make([]int, 0)
	source := make(map[int]int, 0)

	for i, v := range numbers {
		source[v] =i
	}

	for i, v := range numbers {
		if _, ok := source[target-v]; ok && source[target-v] != i {
			out = append(out, i, source[target-v])
			return out
		}
	}

	return out
}