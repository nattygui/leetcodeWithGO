package canmake

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != arr[i-1]-arr[i-2] {
			return false
		}
	}
	return true
}
