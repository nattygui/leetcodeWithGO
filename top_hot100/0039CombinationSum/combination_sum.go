package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	walk(candidates, target, 0, []int{}, &result)
	return result
}

func walk(candidates []int, target int, sum int, curSlice []int, result *[][]int) {
	for i := 0; i < len(candidates); i++ {
		candidate := candidates[i]
		switch {
		case sum+candidate == target:
			oneResult := make([]int, len(curSlice))
			copy(oneResult, curSlice)
			*result = append(*result, append(oneResult, candidate))
			break
		case sum+candidate < target:
			if sum+candidate*2 <= target {
				walk(candidates[i:], target, sum+candidate, append(curSlice, candidate), result)
			} else {
				walk(candidates[i+1:], target, sum+candidate, append(curSlice, candidate), result)
			}
		default:
			break
		}
	}
}
