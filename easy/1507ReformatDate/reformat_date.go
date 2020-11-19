package reformatdate

import (
	"fmt"
	"strconv"
	"strings"
)
func reformatDate(date string) string {
	dateFormat := "%d-%d%d-%d%d"

	dateStr := strings.Split(date, " ")
	year, _ := strconv.Atoi(dateStr[2])
	month := getMonth(dateStr[1])
	day := getDay(dateStr[0])

    month1 := 0
    month2 := month
    if month > 10 {
        month1 = month / 10
        month2 = month % 10
    }

    day1 := 0
    day2 := day
    if day > 10 {
        day1 = day / 10
        day2 = day % 10
    }

	return fmt.Sprintf(dateFormat, year, month1, month2, day1, day2)
}

func getMonth(month string) int {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for i, v := range months {
		if month == v {
			return i + 1
		}
	}
	return 1
}

func getDay(day string) int {
	end := 0
	for i := 0; i < len(day); i++ {
		if day[i] < '0' || day[i] > '9' {
			end = i
			break
		}
	}
	dayNum := day[:end]
	res, _ := strconv.Atoi(dayNum)
	return res
}
