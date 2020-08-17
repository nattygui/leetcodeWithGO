package selfdividingnumbers

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for left <= right {
		if isSelfDividing(left) {
			res = append(res, left)
		}
		left++
	}
	return res
}

func isSelfDividing(num int) bool {
	temp := num
	for temp != 0 {
		if temp%10 == 0 || num%(temp%10) != 0 {
			return false
		}
		temp /= 10
	}
	return true
}
