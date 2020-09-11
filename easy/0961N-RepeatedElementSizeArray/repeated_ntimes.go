package ntimes

func repeatedNTimes(A []int) int {
	valueMap := make(map[int]bool)

	for _, v := range A {
		if !valueMap[v] {
			valueMap[v] = true
			continue
		}
		return v
	}
	return 0
}
