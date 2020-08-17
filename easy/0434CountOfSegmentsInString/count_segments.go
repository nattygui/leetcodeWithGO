package _434CountOfSegmentsInString

func countSegments(s string) int {
	if s == "" {
		return 0
	}

	count := 0

	start := 0
	end := len(s) - 1

	for start < len(s) && end >= 0 && start < end {
		if s[start] != ' ' && s[end] != ' ' {
			break
		}
		if s[start] == ' ' {
			start++
		}
		if s[end] == ' ' {
			end--
		}
	}

	if  s[start] == ' ' && s[end] == ' ' {
		return 0
	}

	for start < end {
		if s[start] == ' ' && s[start+1] != ' ' {
			count++
		}
		start++
	}

	return count+1
}
