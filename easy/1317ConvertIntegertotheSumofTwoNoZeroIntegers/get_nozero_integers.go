package getnozerointegers

func getNoZeroIntegers(n int) []int {
	one := 1
	two := n - one
	for one <= two {
		if !hasZeroInteger(one) && !hasZeroInteger(two) {
			return []int{one, two}
		}
		one++
		two--
	}
	return []int{}
}

func hasZeroInteger(num int) bool {
	for num != 0 {
		if num%10 == 0 {
			return true
		}
		num /= 10
	}
	return false
}
