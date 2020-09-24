package canthreepartsequalsum

func canThreePartsEqualSum(A []int) bool {
	sums := sum(A)
	evenSum := sums / 3
	if sums%3 != 0 {
		return false
	}

	frontIndex := 0
	frontSum := 0
	for frontIndex < len(A) {
		frontSum += A[frontIndex]
		if frontSum == evenSum {
			break
		}
		frontIndex++
	}

	tailIndex := len(A) - 1
	tailSum := 0
	for tailIndex >= 0 {
		tailSum += A[tailIndex]
		if tailSum == evenSum {
			break
		}
		tailIndex--
	}

	if frontSum != evenSum || tailSum != evenSum {
		return false
	}

	if (tailIndex - frontIndex) <= 1 {
		return false
	}

	frontIndex++
	betweenSum := 0
	for frontIndex < tailIndex {
		betweenSum += A[frontIndex]
		frontIndex++
	}

	if betweenSum != evenSum {
		return false
	}

	return true
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
