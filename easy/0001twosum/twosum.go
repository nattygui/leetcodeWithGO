package twosum

func twoSum(nums []int, target int) []int {
	out := make([]int, 0)
	source := make(map[int]int, len(nums))

	for i, v := range nums {
		source[v] = i
	}

	for i, v := range nums {
		if _, ok := source[target-v]; ok && source[target-v] != i {
			out = append(out, i, source[target-v])
			return out
		}
	}
	return out
}
