package monotonic

func isMonotonic(A []int) bool {
	length := len(A)
	flag := 0
	for i := 1; i < length; i++ {
		temp := A[i] - A[i-1]
		switch {
		case temp > 0:
			if flag == 0 {
				flag = 1
				continue
			}

			if temp*flag < 0 {
				return false
			}

		case temp == 0:
			continue
		case temp < 0:
			if flag == 0 {
				flag = -1
				continue
			}

			if temp*flag < 0 {
				return false
			}
		}
	}
	return true
}
