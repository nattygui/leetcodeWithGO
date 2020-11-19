package countgoodtriplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for z := j + 1; z < length; z++ {
				if isGoodTriplets(arr[i], arr[j], arr[z], a, b, c) {
					res++
				}
			}
		}
	}
	return res
}

func isGoodTriplets(i, j, k int, a, b, c int) bool {
	if i-j < -a || i-j > a {
		return false
	}
	if j-k < -b || j-k > b {
		return false
	}
	if i-k < -c || i-k > c {
		return false
	}
	return true
}
