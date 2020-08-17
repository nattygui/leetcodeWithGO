package _509FibonacciNumber

func fib(N int) int {
	if N <= 1 {
		return N
	}
	bp := make([]int, N+1)

	bp[0] = 0
	bp[1] = 1

	for i := 2; i < N+1; i++ {
		bp[i] = bp[i-1] + bp[i-2]
	}

	return bp[N]

}