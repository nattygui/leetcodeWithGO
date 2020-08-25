package countprime

func countPrimeSetBits(L int, R int) int {
	res := 0
	for L <= R {
		bits := get1bits(L)
		if isPrimeNumber(bits) {
			res++
		}
		L++
	}
	return res
}

func get1bits(value int) int {
	res := 0
	for value != 0 {
		if value%2 == 1 {
			res++
		}
		value /= 2
	}
	return res
}

func isPrimeNumber(value int) bool {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	for _, prime := range primes {
		if prime == value {
			return true
		}
	}
	return false
}
