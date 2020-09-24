package addarray

func addToArrayForm(A []int, K int) []int {
	res := make([]int, 0)
	carry := 0
	curSum := 0
	curValue := 0
	index := len(A) - 1
	indexValue := 0
	for {
		if index < 0 && K == 0 && carry == 0 {
			break
		}

		if index < 0 {
			indexValue = 0
		} else {
			indexValue = A[index]
		}

		curSum = indexValue + K%10 + carry
		curValue = curSum % 10
		carry = curSum / 10
		res = append(res, curValue)
		index--
		K /= 10
	}

	reverseArray(res)
	return res
}

func reverseArray(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}
}
