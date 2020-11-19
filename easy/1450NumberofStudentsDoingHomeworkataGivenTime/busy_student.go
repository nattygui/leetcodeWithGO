package busystudent

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	result := 0
	students := len(startTime)
	for i := 0; i < students; i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			result++
		}
	}
	return result
}
