package _069Sqrt

func mySqrt(x int) int {
	left, right := 0, x
	if x <= 1 {
		return x
	}
	mid := 0
	for left < right {
		mid = left + (right - left) / 2
		if mid * mid == x {
			return mid
		}
		if mid*mid > x {
			right = mid
		}
		if mid*mid < x {
			left = mid
		}
		if mid*mid < x && (mid+1)*(mid+1) > x{
			return mid
		}
	}
	return 0
}
