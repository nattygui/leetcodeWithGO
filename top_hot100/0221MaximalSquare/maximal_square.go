package maximalsquare

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
