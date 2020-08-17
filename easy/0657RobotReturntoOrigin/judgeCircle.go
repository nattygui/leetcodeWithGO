package _657RobotReturntoOrigin

func judgeCircle(moves string) bool {
	record := []int{0, 0}
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'R':
			record[0]--
		case 'L':
			record[0]++
		case 'U':
			record[1]++
		case 'D':
			record[1]--
		}
	}
	return record[0] ==0 && record[1] == 0
}
