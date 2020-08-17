package _009PalindromeNumber

func palindromeNumber(x int) bool {
	t := x
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = res * 10 + temp
		x = x / 10
	}

	if res != t {
		return false
	}

	return true
}
