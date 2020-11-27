# 64. 最小路径和

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例 1：

```
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
```

示例 2：

```
输入：grid = [[1,2,3],[4,5,6]]
输出：12
```

## 解法一

递归 超时

```go
func minPathSum(grid [][]int) int {
	rows := len(grid)-1
	cols := len(grid[0])-1
	result := math.MaxInt64
	walk(grid, rows, cols, 0, &result)
	return result
}

func walk(grid [][]int, m, n, curscore int, result *int) {
	curscore += grid[m][n]
	if m == 0 && n == 0 {
		if *result > curscore {
			*result = curscore
		}
		return
	}
	if m > 0 {
		walk(grid, m-1, n, curscore, result)
	}
	if n > 0 {
		walk(grid, m, n-1, curscore, result)
	}
}
```

## 解法二

- 时间复杂度： O(mxn)

- 空间复杂度: O(1)

使用动态规划，当前的节点，只有俩种路径，左边或者上边。

只要获取最小的那一个值就ok了， 状态转移方程： grid[i][j] += min(grid[i][j-1], grid[i-1][j])

```go
func minPathSum(grid [][]int) int {
	// 初始化
	rows := len(grid)
	cols := len(grid[0])
	for i := 1; i < rows; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < cols; j++ {
		grid[0][j] += grid[0][j-1]
	}
	// dp 
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[rows-1][cols-1]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
```