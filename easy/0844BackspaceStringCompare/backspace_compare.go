package backspacecompare

func backspaceCompare(S string, T string) bool {
	news := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if S[i] != '#' {
			news = append(news, S[i])
			continue
		}
		if len(news) == 0 {
			continue
		}
		news = news[:len(news)-1]
	}
	newt := make([]byte, 0)
	for i := 0; i < len(T); i++ {
		if T[i] != '#' {
			newt = append(newt, T[i])
			continue
		}
		if len(newt) == 0 {
			continue
		}
		newt = newt[:len(newt)-1]
	}

	if string(newt) == string(news) {
		return true
	}
	return false
}
