# 200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

 

示例 1：
```
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
```

## 解法 

遍历每个值，若当前值为 ‘1’ 则说明存在一个岛屿。将该岛屿所有的值变成‘2’。继续寻找下一个岛屿。

```go
var directs [4][2]int = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func numIslands(grid [][]byte) int {
	result := 0
	for x, row := range grid {
		for y, num := range row {
			if num == '2' {
				continue
			}
			if num == '1' {
				grid[x][y] = '2'
				findAll(x, y, grid)
				result++
			}
		}
	}
	return result
}

func findAll(x, y int, grid [][]byte) {
	for _, direct := range directs {
		newx := x + direct[0]
		if newx < 0 || newx >= len(grid) {
			continue
		}
		newy := y + direct[1]
		if newy < 0 || newy >= len(grid[0]) {
			continue
		}
		if grid[newx][newy] == '1' {
			grid[newx][newy] = '2'
			findAll(newx, newy, grid)
		}
	}
}
```