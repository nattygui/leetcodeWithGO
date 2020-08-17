package _643MaximumAverageSubarrayI

func findMaxAverage(nums []int, k int) float64 {
	i, j := 0, k-1
	max := 0
	temp := 0
	for s := 0; s <= j; s++ {
		max += nums[s]
	}
	temp = max

	for j+1 < len(nums) {
		temp = temp - nums[i] + nums[j+1]
		if temp > max {
			max = temp
		}
		i++
		j++
	}
	return float64(max) / float64(k)
}
