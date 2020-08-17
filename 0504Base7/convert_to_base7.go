package _504Base7

import "strconv"

func convertToBase7(num int) string {
	resint := 0
	count := 1

	for num != 0 {
		resint = num % 7 * count
		count *= 10
		num /= 7
	}

	return strconv.Itoa(resint)

}