package bitwisecomplement

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	result := 0
	carry := 1
	for N > 0 {
		result += reverseBit(N%2) * carry
		carry *= 2
		N /= 2
	}
	return result
}

func reverseBit(num int) int {
	if num == 0 {
		return 1
	}
	return 0
}
