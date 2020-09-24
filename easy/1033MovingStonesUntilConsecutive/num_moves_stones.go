package nummovesstones

func numMovesStones(a int, b int, c int) []int {
	min, max := a, c
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	if max < a {
		max = a
	}
	if max < b {
		max = b
	}

	medium := a + b + c - min - max
	minMove := 0
	maxMove := max - min - 2
	if maxMove == 0 {
		minMove = 0
	} else if max-medium > 2 && medium-min > 2 {
		minMove = 2
	} else {
		minMove = 1
	}
	return []int{minMove, maxMove}
}
