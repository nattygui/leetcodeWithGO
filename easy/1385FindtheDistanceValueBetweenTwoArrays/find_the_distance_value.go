package findthedistancevalue

import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	result := 0
	for _, num := range arr1 {
		if isPerfectNum(num, arr2, d) {
			result++
		}
	}
	return result
}

func isPerfectNum(num int, arr []int, d int) bool {
	for _, a := range arr {
		if num <= a+d && num >= a-d {
			return false
		}
	}
	return true
}
