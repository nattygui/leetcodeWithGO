package largestsumafterknegations

import (
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)

	index := 0
	for K > 0 {
		switch {
		case A[index] < 0:
			A[index] *= -1
			index++
			if index >= len(A) {
				index = len(A) - 1
			}
			K--
		case A[index] > 0:
			curIndex := index
			if index-1 >= 0 && A[index-1] < A[curIndex] {
				curIndex = index - 1
			}
			if index+1 < len(A) && A[index+1] < A[curIndex] {
				curIndex = index + 1
			}
			A[curIndex] *= -1
			K--
		default:
			K--
		}
	}

	return sum(A)
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
