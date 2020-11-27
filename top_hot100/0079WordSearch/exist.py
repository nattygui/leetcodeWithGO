# coding: utf-8
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