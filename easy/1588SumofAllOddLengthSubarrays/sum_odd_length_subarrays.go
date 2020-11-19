package sumoddlengthsubarrays

func sumOddLengthSubarrays(arr []int) int {
	length := len(arr)
	sum := 0
	all := 1
	for all <= length {
		for i := 0; i < length-all+1; i++ {
			for j := i; j < i+all; j++ {
				sum += arr[j]
			}
		}
		all += 2
	}
	return sum
}
