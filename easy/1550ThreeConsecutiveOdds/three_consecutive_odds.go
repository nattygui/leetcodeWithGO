package threeconsecutiveodds

func threeConsecutiveOdds(arr []int) bool {
	result := false
	for i := 0; i < len(arr)-2; i++ {
		if !isOdd(arr[i]) {
			continue
		}
		if !isOdd(arr[i+1]) {
			i++
			continue
		}
		if isOdd(arr[i+2]) {
			result = true
		}
	}
	return result
}

func isOdd(n int) bool {
	return n%2 == 1
}
