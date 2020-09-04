package lemonadechange

func lemonadeChange(bills []int) bool {
	smallChangeMap := map[int]int{
		5:  0,
		10: 0,
	}

	for _, bill := range bills {
		if bill == 5 {
			smallChangeMap[5]++
			continue
		}
		if bill == 10 {
			if smallChangeMap[5] != 0 {
				smallChangeMap[5]--
				smallChangeMap[10]++
				continue
			}
			return false
		}
		if bill == 20 {
			if smallChangeMap[10] > 0 && smallChangeMap[5] > 0 {
				smallChangeMap[5]--
				smallChangeMap[10]--
				continue
			}

			if smallChangeMap[5] > 2 {
				smallChangeMap[5] -= 3
				continue
			}

			return false
		}
	}
	return true
}
