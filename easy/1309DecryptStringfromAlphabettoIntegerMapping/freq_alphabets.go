package freqalphabets

import "strconv"

func freqAlphabets(s string) string {
	// 首先获取有多少个#
	spcialSymbolnum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			spcialSymbolnum++
		}
	}
	spcialSymbolnum = spcialSymbolnum + len(s) - 3*spcialSymbolnum

	resultBytes := make([]byte, spcialSymbolnum)
	spcialSymbolnum--
	benchChar := byte('a' - 1)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			tempInt, _ := strconv.Atoi(s[i-2 : i])
			resultBytes[spcialSymbolnum] = benchChar + byte(tempInt)
			i = i - 2
			spcialSymbolnum--
			continue
		}
		temp := s[i] - byte('0')
		resultBytes[spcialSymbolnum] = temp + benchChar
		spcialSymbolnum--
	}

	return string(resultBytes)
}
