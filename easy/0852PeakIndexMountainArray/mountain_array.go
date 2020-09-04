package peakindex

func peakIndexInMountainArray(A []int) int {
	i := 0
	j := len(A) - 1
	for i < j {
		if A[i] > A[j] {
			j--
			continue
		}
		i++
	}

	return i
}
