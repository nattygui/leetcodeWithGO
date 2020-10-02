package minimumabsdifference

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	result := make([][]int, 0)
	if len(arr) == 0 {
		return result
	}
	minDiff := math.MaxInt64
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if minDiff > diff {
			minDiff = diff
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] == arr[i]+minDiff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}
