package maxpower

func maxPower(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := 1
	temp := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			temp++
		} else {
			temp = 1
		}
		if temp != 1 && max < temp {
			max = temp
		}
	}
	return max
}
