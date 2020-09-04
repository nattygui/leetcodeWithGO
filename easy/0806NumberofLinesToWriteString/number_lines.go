package numberoflines

func numberOfLines(widths []int, S string) []int {
	lines := 0
	units := 0
	for _, letter := range S {
		temp := widths[int(letter-'a')]
		if units+temp < 100 {
			units += temp
		} else if units+temp == 100 {
			units = 0
			lines++
		} else if units+temp > 100 {
			units = temp
			lines++
		}
	}
	return []int{lines + 1, units}
}
