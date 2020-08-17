package _204CountPrimes

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	isPrimes := make([]bool, n)
	count := 1
	for j := 2; (2 * j) < n; j++ {
		isPrimes[2*j] = true
	}

	for i := 3; i < len(isPrimes); i=i+2 {
		if !(isPrimes[i]) {
			for j := 2; (i * j) < n; j++ {
				isPrimes[i*j] = true
			}
			count++
		}
	}
	return count
}
