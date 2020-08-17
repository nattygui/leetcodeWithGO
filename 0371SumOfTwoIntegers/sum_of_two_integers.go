package _371SumOfTwoIntegers

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	lower := 0
	high := 0

	for true {
		lower = a ^ b
		high = a & b
		if high == 0 {
			break
		}

		a = lower
		b = high << 1
	}
	return lower
}
