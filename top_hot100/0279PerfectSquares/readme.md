# 279. 完全平方数
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:
```
输入: n = 12
输出: 3 
解释: 12 = 4 + 4 + 4.
```
示例 2:
```
输入: n = 13
输出: 2
解释: 13 = 4 + 9.
```

## 解法

```go
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
```