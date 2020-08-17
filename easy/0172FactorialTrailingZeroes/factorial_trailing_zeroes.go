package _172FactorialTrailingZeroes

func trailingZeroes(n int) int {
	fiveNum := 0
	for n >= 5 {
		fiveNum += n / 5
		n /= 5
	}
	return fiveNum
}
