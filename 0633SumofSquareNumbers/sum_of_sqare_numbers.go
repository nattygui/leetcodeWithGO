package _633SumofSquareNumbers

// 二分查找法
func judgeSquareSum(c int) bool {
	if c == 0 || c == 1 || c == 2 {
		return true
	}
	i := 0
	for c-i*i >= 0 {
		if search(c - i*i) {
			return true
		}
		i++
	}
	return false
}

func search(num int) bool {
	i := 0
	j := num
	for i < j {
		mid := (i + j) / 2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return false
}
