package _485MaxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	nums = append(nums, 0)

	max := 0
	i := 0

	for j, v := range nums {
		if v != 1 {
			temp := j - i
			if max < temp {
				max = temp
			}
			i = j + 1
		}
	}

	return max
}
