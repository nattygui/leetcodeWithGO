package findkthpositive

func findKthPositive(arr []int, k int) int {
	for i := range arr {
		if arr[i] <= k {
			k++
		}
	}
	return k
}
