package slowestkey

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var result byte
	maxTime := -1
	for i := 0; i < len(releaseTimes); i++ {
		if i == 0 {
			maxTime = releaseTimes[i]
			result = keysPressed[i]
			continue
		}
		diff := releaseTimes[i] - releaseTimes[i-1]
		if diff > maxTime {
			maxTime = diff
			result = keysPressed[i]
		} else if diff == maxTime && result < keysPressed[i] {
			result = keysPressed[i]
		}
	}
	return result
}
