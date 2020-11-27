# 79. 单词搜索

给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

```
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
```

# 解法

通过递归和回溯判断

golang
```go
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	rows := len(board)
	cols := len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				board[i][j] = '0'
				if findOnechar(board, i, j, 1, word) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}
	return false
}

func findOnechar(board [][]byte, i, j, num int, word string) bool {
	if len(word) == num {
		return true
	}
	curChar := word[num]
	if i-1 >= 0 && i-1 < len(board) {
		if board[i-1][j] == curChar {
			board[i-1][j] = '0'
			if findOnechar(board, i-1, j, num+1, word) {
				return true
			}
			board[i-1][j] = curChar
		}
	}
	if i+1 >= 0 && i+1 < len(board) {
		if board[i+1][j] == curChar {
			board[i+1][j] = '0'
			if findOnechar(board, i+1, j, num+1, word) {
				return true
			}
			board[i+1][j] = curChar
		}
	}
	if j-1 >= 0 && j-1 < len(board[0]) {
		if board[i][j-1] == curChar {
			board[i][j-1] = '0'
			if findOnechar(board, i, j-1, num+1, word) {
				return true
			}
			board[i][j-1] = curChar
		}
	}
	if j+1 >= 0 && j+1 < len(board[0]) {
		if board[i][j+1] == curChar {
			board[i][j+1] = '0'
			if findOnechar(board, i, j+1, num+1, word) {
				return true
			}
			board[i][j+1] = curChar
		}
	}
	return false
}
```

python
```python
class Solution:
    def exist(self, board, word):
        if len(word) == 0:
            return True
        rows = len(board)
        cols = len(board[0])
        for i in range(rows):
            for j in range(cols):
                if board[i][j] == word[0]:
                    board[i][j] = '0'
                    if self.find_char(board, i, j, 1, word) is True:
                        return True
                    board[i][j] = word[0]
        
        return False
    
    def find_char(self, board, i, j, num, word):
        if len(word) == num:
            return True
        
        if i-1 >= 0 and i-1 < len(board):
            if board[i-1][j] is word[num]:
                board[i-1][j] = '0'
                if self.find_char(board, i-1, j, num+1, word) is True:
                    return True
                board[i-1][j] = word[num]

        if i+1 >= 0 and i+1 < len(board):
            if board[i+1][j] is word[num]:
                board[i+1][j] = '0'
                if self.find_char(board, i+1, j, num+1, word) is True:
                    return True
                board[i+1][j] = word[num]
            
        if j-1 >= 0 and j-1 < len(board[0]):
            if board[i][j-1] is word[num]:
                board[i][j-1] = '0'
                if self.find_char(board, i, j-1, num+1, word) is True:
                    return True
                board[i][j-1] = word[num]

        if j+1 >= 0 and j+1 < len(board[0]):
            if board[i][j+1] is word[num]:
                board[i][j+1] = '0'
                if self.find_char(board, i, j+1, num+1, word) is True:
                    return True
                board[i][j+1] = word[num]

        return False
```