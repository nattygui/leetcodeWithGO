package _645SetMismatch

func findErrorNums(nums []int) []int {
	staticNumsMap := make(map[int]bool, 0)
	sum := (1 + len(nums)) * len(nums) / 2
	res := []int{}
	for _, v := range nums {
		if _, ok := staticNumsMap[v]; ok {
			res = append(res, v)
		} else {
			staticNumsMap[v] = true
			sum -= v
		}
	}
	return append(res, sum)
}
