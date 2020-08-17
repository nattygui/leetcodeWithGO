package _367ValidPerfectSquare

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left := 0
	right := num / 2
	mid := 0
	temp := 0
	for left < right {
		mid = (left + right) / 2
		temp = mid * mid

		if temp == num {
			return true
		} else if temp < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false
}
