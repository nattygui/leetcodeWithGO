# 62. 不同路径

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？

示例 1:

```
输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
```

示例 2:

```
输入: m = 7, n = 3
输出: 28
```

## 解法一

递归，一直递减，直到到达最后目标

```go
func uniquePaths(m int, n int) int {
	result := 0
	walk(m-1, n-1, &result)
	return result
}

func walk(m, n int, result *int) {
	if m == 0 && n == 0 {
		*result = *result + 1
		return
	}
	if m >= 0 {
		walk(m-1, n, result)
	}
	if n >= 0 {
		walk(m, n-1, result)
	}
}
```

## 解法二

通过动态规划来解题，状态转移方程 dp[i][j] = dp[i-1][j] + dp[i][j-1]

```go
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	// dp 初始化
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	// 开始计算
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
            // 状态转移
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
```
