package _461HammingDistance

func hammingDistance(x int, y int) int {
	getMaxbit := func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		count := 0
		for n != 0 {
			if n % 2 == 1 {
				count++
			}
			n /= 2
		}
		return count
	}

	num := x^y

	return getMaxbit(num)
}
