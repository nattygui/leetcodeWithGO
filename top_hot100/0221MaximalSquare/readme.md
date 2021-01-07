# 221. 最大正方形
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

示例 1：

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4
示例 2：

输入：matrix = [["0","1"],["1","0"]]
输出：1

## 解法

那么如何计算 dp 中的每个元素值呢？对于每个位置 (i, j)，检查在矩阵中该位置的值：

如果该位置的值是 0，则 dp(i, j) = 0，因为当前位置不可能在由 1 组成的正方形中；

如果该位置的值是 1，则 dp(i, j)的值由其上方、左方和左上方的三个相邻位置的 dp 值决定。具体而言，当前位置的元素值等于三个相邻位置的元素中的最小值加 1，状态转移方程如下：

dp(i, j)=min(dp(i−1, j), dp(i−1, j−1), dp(i, j−1))+1

```go
func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	if rows == 1 && cols == 1 {
		if matrix[0][0] == '0' {
			return 0
		}
		return 1
	}
	result := 0
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
		if matrix[i][0] == '0' {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
			result = 1
		}
	}
	for j := 1; j < cols; j++ {
		if matrix[0][j] == '0' {
			dp[0][j] = 0
		} else {
			dp[0][j] = 1
			result = 1
		}
	}
	// 求值
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			minV := min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			dp[i][j] = minV + 1
			if result < minV+1 {
				result = minV + 1
			}
		}
	}
	return result * result
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
```