package _566ReshapeMatrix

func matrixReshape(nums [][]int, r int, c int) [][]int {
	l, w := len(nums), len(nums[0])
	if r * c != l * w {
		return nums
	}

	res := make([][]int, r)

	for i := 0; i < r; i++ {
		temp := make([]int, c)
		for j := 0; j < c; j++ {
			sum := i * c + j
			temp[j] = nums[sum / w][sum % w]
		}
		res[i] = temp
	}
	return res
}