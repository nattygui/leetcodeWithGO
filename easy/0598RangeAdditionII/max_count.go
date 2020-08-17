package _598RangeAdditionII

import "math"

func maxCount(m int, n int, ops [][]int) int {
	minx, miny := math.MaxInt64, math.MaxInt64
	for _, v := range ops {
		if  minx > v[0] {
			minx = v[0]
		}
		if  miny > v[1] {
			miny = v[1]
		}
	}
	if minx > m {
		minx = m
	}
	if miny > n {
		miny = n
	}
	return minx * miny
}
