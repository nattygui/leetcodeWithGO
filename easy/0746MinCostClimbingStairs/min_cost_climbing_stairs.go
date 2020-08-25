package mincost

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	size := len(cost)
	dp := make([]int, size)
	dp[0] = 0
	dp[1] = min(cost[0], cost[1])

	for i := 2; i < size; i++ {
		dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i-1])
	}

	return dp[size-1]
}

func min(value1, value2 int) int {
	if value1 > value2 {
		return value2
	}
	return value1
}
