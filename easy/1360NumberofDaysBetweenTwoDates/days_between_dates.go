package daysbetweendates

import (
	"strconv"
	"strings"
)

var months = [...]int{
	0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func daysBetweenDates(date1 string, date2 string) int {
	result := getDays(date2) - getDays(date1)
	if result >= 0 {
		return result
	}
	return -result
}

func getDays(date string) int {
	ymd := strings.Split(date, "-")
	year, _ := strconv.Atoi(ymd[0])
	month, _ := strconv.Atoi(ymd[1])
	day, _ := strconv.Atoi(ymd[2])

	days := 0
	for y := 1971; y < year; y++ {
		days += isLeap(y) + 365
	}

	for m := 1; m < month; m++ {
		days += months[m]
	}

	days += isLeap(year)
	return days + day
}

func isLeap(year int) int {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return 1
	}
	return 0
}
