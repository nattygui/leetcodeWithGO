package validmountainarray

func validMountainArray(A []int) bool {
	firstIndex, lastIndex := -1, -1
	for i := 1; i < len(A); i++ {
		if A[i]-A[i-1] <= 0 {
			break
		}
		firstIndex = i
	}
	for i := len(A) - 2; i >= 0; i-- {
		if A[i+1]-A[i] >= 0 {
			break
		}
		lastIndex = i
	}

	if firstIndex != lastIndex || (firstIndex == -1 && lastIndex == -1) {
		return false
	}
	return true
}
