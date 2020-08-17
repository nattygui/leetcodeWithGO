package _171ExcelSheetColumnNumber

func titleToNumber(s string) int {
	sum := 0
	temp := 1
	for i := len(s)-1; i >= 0; i-- {
		sum += int(s[i] - 'A' + 1) * temp
		temp *= 26
	}
	return sum
}