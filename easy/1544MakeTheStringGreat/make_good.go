package makegood

func makeGood(s string) string {
	results := []byte(s)
	findFlag := true
	for findFlag {
		findFlag = false

		for i := 0; i < len(results)-1; i++ {
			if isGood(results[i], results[i+1]) {
				results = append(results[:i], results[i+2:]...)
				findFlag = true
			}
		}
	}
	return string(results)
}

func isGood(x, y byte) bool {
	if x < y {
		x, y = y, x
	}
	if x-y == 32 {
		return true
	}
	return false
}
