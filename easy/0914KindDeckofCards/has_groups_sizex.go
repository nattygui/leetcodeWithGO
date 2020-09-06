package hasgroupssizex

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 0 {
		return false
	}
	m := make(map[int]int)
	for _, v := range deck {
		m[v]++
	}
	maxVal := m[deck[0]]
	for _, v := range m {
		maxVal = gcd(v, maxVal)
		if maxVal < 2 {
			return false
		}
	}
	return true
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	}

	return y
}
