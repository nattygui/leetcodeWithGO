package countlargestgroup

func countLargestGroup(n int) int {
	numArray := map[int]int{}

	for i := 1; i <= n; i++ {
		num := 0
		label := i
		for label != 0 {
			num = num + label%10
			label = label / 10
		}

		numArray[num]++
	}

	maxValue := 0
	count := 0
	for _, v := range numArray {
		if v > maxValue {
			maxValue = v
			count = 1
		} else if v == maxValue {
			count++
		}
	}

	return count
}
