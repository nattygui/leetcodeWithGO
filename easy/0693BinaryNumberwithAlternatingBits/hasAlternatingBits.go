package _693BinaryNumberwithAlternatingBits

func hasAlternatingBits(n int) bool {
	res := true
	pre := n % 2
	n = n / 2
	for n != 0 {
		if n%2 != pre {
			pre = n % 2
		} else {
			return false
		}

		n = n / 2
	}

	return res
}
