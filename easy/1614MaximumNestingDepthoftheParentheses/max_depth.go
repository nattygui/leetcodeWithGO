package maxdepth

func maxDepth(s string) int {
	var pair, max int
	for _, v := range s {
		// 当前的原语字符串
		switch v {
		case '(':
			pair++
		case ')':
			if pair > max {
				max = pair
			}
			pair--
		}
	}
	return max
}
