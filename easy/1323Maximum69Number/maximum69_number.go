package maximum69number

func maximum69Number(num int) int {
	digits := getMax6numberDigits(num)
	if digits == 0 {
		return num
	}
	return num + digits*3
}

func getMax6numberDigits(num int) int {
	digits := 1
	result := 0
	for num != 0 {
		if num%10 == 6 {
			result = digits
		}
		num /= 10
		digits *= 10
	}
	return result
}
