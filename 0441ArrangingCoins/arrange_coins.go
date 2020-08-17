package _441ArrangingCoins

func arrangeCoins(n int) int {
	for i := 1; i < n; i++ {
		left := (i * i + i) / 2
		right := (((i+1) * (i+1)) + i + 1) / 2
		if left <= n && n < right {
			return i
		}
	}
	return 0
}
