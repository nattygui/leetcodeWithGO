package tictactioe

func tictactoe(moves [][]int) string {
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
	}

	for index, move := range moves {
		if index%2 == 0 {
			matrix[move[0]][move[1]] = 1
		} else {
			matrix[move[0]][move[1]] = -1
		}

		if index <= 2 {
			continue
		}

		if temp := check(matrix); temp != "" {
			return temp
		}
	}
	if checkHasEmpty(matrix) {
		return "Pending"
	}

	return "Draw"
}

func checkHasEmpty(martix [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if martix[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func check(martix [][]int) string {
	for i := 0; i < 3; i++ {
		if temp := martix[i][0] + martix[i][1] + martix[i][2]; temp == 3 {
			return "A"
		} else if temp == -3 {
			return "B"
		}

		if temp := martix[0][i] + martix[1][i] + martix[2][i]; temp == 3 {
			return "A"
		} else if temp == -3 {
			return "B"
		}
	}

	if temp := martix[0][0] + martix[1][1] + martix[2][2]; temp == 3 {
		return "A"
	} else if temp == -3 {
		return "B"
	}

	if temp := martix[0][2] + martix[1][1] + martix[2][0]; temp == 3 {
		return "A"
	} else if temp == -3 {
		return "B"
	}

	return ""
}
