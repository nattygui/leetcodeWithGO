package numprimearrangements

import "math"

var mod = 1000000007

func numPrimeArrangements(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}
	if n == 1 {
		return 1
	}
	tmpPrime := 1
	for i := 1; i <= count; i++ {
		tmpPrime = tmpPrime * i % mod
	}
	tmpNotPrime := 1
	for i := 1; i <= n-count; i++ {
		tmpNotPrime = tmpNotPrime * i % mod
	}
	if count == 0 {
		return tmpNotPrime
	}
	if count == n {
		return tmpPrime
	}
	return tmpPrime * tmpNotPrime % mod
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
