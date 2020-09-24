package numrookcaptures

type position struct {
	x int
	y int
}

func numRookCaptures(board [][]byte) int {
	pos := &position{}
	find := pos.findRChest(board)
	if !find {
		return 0
	}

	return pos.findAllChest(board)
}

func (p *position) findRChest(board [][]byte) bool {
	for axisy, chesty := range board {
		for axisx, chestx := range chesty {
			if chestx == 'R' {
				p.x = axisx
				p.y = axisy
				return true
			}
		}
	}
	return false
}

func (p position) findAllChest(board [][]byte) int {
	res := 0
	tempy := p.y + 1
	for tempy < 8 {
		times, find := isCorrectChest(board, p.x, tempy)
		if find {
			res += times
			break
		}
		tempy++
	}

	tempy = p.y - 1
	for tempy >= 0 {
		times, find := isCorrectChest(board, p.x, tempy)
		if find {
			res += times
			break
		}
		tempy--
	}

	tempx := p.x + 1
	for tempx < 8 {
		times, find := isCorrectChest(board, tempx, p.y)
		if find {
			res += times
			break
		}
		tempx++
	}

	tempx = p.x - 1
	for tempx >= 0 {
		times, find := isCorrectChest(board, tempx, p.y)
		if find {
			res += times
			break
		}
		tempx--
	}
	return res
}

func isCorrectChest(board [][]byte, x, y int) (int, bool) {
	find := false
	res := 0
	if board[y][x] == 'B' {
		find = true
	}
	if board[y][x] == 'p' {
		res++
		find = true
	}
	return res, find
}
