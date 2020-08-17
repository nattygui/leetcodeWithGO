package _697DegreeofanArray

import "fmt"

func findShortestSubArray(nums []int) int {

	staticMap := map[int]*statics{}

	for i, v := range nums {
		sis, ok := staticMap[v]
		if ok {
			sis.count++
			sis.length = i - sis.firstIndex
		} else {
			sis = &statics{
				count:      1,
				length:     0,
				firstIndex: i,
			}
		}

		staticMap[v] = sis
	}

	maxDispear := 0
	shortLength := len(nums)

	for key, sis := range staticMap {
		fmt.Println(key, sis.count, sis.length)
		if sis.count > maxDispear {
			maxDispear = sis.count
			shortLength = sis.length
		} else if sis.count == maxDispear && sis.length <= shortLength {
			maxDispear = sis.count
			shortLength = sis.length
		}
	}

	return shortLength + 1
}

type statics struct {
	count      int
	length     int
	firstIndex int
}
