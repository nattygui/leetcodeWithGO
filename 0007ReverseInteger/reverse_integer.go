package _007ReverseInteger

import "math"

func reverseInteger(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = res * 10 + temp
		x = x / 10
	}

	res = sign * res

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}