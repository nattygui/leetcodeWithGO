package _521LongestUncommonSubsequenceI

func findLUSlength(a string, b string) int {
	x := len(a)
	y := len(b)

	if x != y {
		return max(x, y)
	}

	for i := 0; i < x; i++ {
		if a[i] != b[i] {
			return x
		}
	}

	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}