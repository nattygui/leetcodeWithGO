# 240. 搜索二维矩阵 II
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

## 解法 

从最后一行第一列开始搜索（i, j)

若当前值大于target，则i--
若当前值小于target，则j++
若当前值等于target，则return true

若i，j 超过了其中一个索引，则返回false

```go
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	i := rows - 1
	j := 0
	for {
		if i < 0 || j > cols-1 {
			return false
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		}
	}
}
```
