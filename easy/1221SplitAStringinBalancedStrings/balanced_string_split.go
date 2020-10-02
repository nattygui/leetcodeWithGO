package balancedstringsplit

func balancedStringSplit(s string) int {
	l := 0
	r := 0
	result := 0
	for _, v := range s {
		if v == 'L' {
			l++
		}
		if v == 'R' {
			r++
		}
		if l == r {
			result++
			l = 0
			r = 0
		}
	}
	return result
}
