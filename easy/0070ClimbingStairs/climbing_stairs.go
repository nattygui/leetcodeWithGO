package _070ClimbingStairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	temp := 0
	l1 := 1
	l2 := 2
	for i := 3; i <= n; i++ {
		temp = l1 + l2
		l1 = l2
		l2 = temp
	}
	return l2
}
