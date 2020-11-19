package trimmean

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	length := len(arr)
	last := length * 5 / 100
	arr = arr[last-1 : length-last]
	return mean(arr)
}

func mean(arr []int) float64 {
	length := len(arr)
	result := 0.0
	for _, num := range arr {
		result += float64(num)
	}
	return result / float64(length)
}
