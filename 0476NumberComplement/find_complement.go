package _476NumberComplement

func findComplement(num int) int {
	maxBit := 0
	s := 0
	for num != 0 {
		s += 1 << maxBit
		maxBit++
		num = num >> 1
	}

	return s - num
}