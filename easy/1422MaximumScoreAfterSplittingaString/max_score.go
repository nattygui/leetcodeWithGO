package maxscore

func maxScore(s string) int {
	zeros := 0
	ones := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	maxRsult := 0
	tempScore := ones
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			tempScore++
		} else {
			tempScore--
		}
		if maxRsult < tempScore {
			maxRsult = tempScore
		}
	}
	return maxRsult
}
