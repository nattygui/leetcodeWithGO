package _551StudentAttendanceRecordI

func checkRecord(s string) bool {
	absent, late := 0, 0

	for i, _ := range s {
		if s[i] == 'A' {
			absent++
		} else if s[i] == 'L' && i + 1 < len(s) && s[i+1] =='L' {
			late++
		} else {
			late = 0
		}
		if absent > 1 || late >= 2 {
			return false
		}
	}

	return true
}
