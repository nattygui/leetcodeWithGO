package _191NumberOf1Bits

func hammingWeight(num uint32) int {
	var temp uint32
	count := 0
	for i := 0; i < 32; i++ {
		temp = num % 2
		num /= 2
		if temp == 1 {
			count++
		}
	}
	return count
}
