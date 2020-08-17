package _448FindAllNumbersDisappearedInAnArray

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	count := 0
	res := make([]int, n)

	for _, v := range nums {
		res[v]++
	}

	for i, v := range res {
		if v == 0 {
			res[count] = i+1
			count++
		}
	}

	return res[0:count]
}
