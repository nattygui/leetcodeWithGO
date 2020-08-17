package _205IsomorphicStrings

func isIsomorphic(s string, t string) bool {
	sBytes := []byte(s)
	tBytes := []byte(t)
	length := len(sBytes)

	smaps := make(map[byte]int, 0)
	scount := 0
	tmaps := make(map[byte]int, 0)
	tcount := 0
	for i := 0; i < length; i++ {
		if _, ok := smaps[sBytes[i]]; ok {
			sBytes[i] = byte(smaps[sBytes[i]])
		} else {
			smaps[sBytes[i]] = scount
			sBytes[i] = byte(scount)
			scount++
		}

		if _, ok := tmaps[tBytes[i]]; ok {
			tBytes[i] = byte(tmaps[tBytes[i]])
		} else {
			tmaps[tBytes[i]] = tcount
			tBytes[i] = byte(tcount)
			tcount++
		}
	}

	if string(sBytes) == string(tBytes) {
		return true
	}
	return false
}
