package exist

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
