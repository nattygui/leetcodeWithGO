package containspattern

func containsPattern(arr []int, m int, k int) bool {
	length := len(arr)
	if length < m {
		return false
	}
	if k == 1 {
		return true
	}
	for i := 0; i < length-1-m; i++ {
		count := 1
		for j := i; j+m < length; j++ {
			if arr[j] != arr[j+m] {
				break
			}
			if (j-i)%m == m-1 {
				count++
			}
		}
		if count >= k {
			return true
		}
	}
	return false
}
