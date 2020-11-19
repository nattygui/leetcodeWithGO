package modifystring

func modifyString(s string) string {
	chars := make([]byte, len(s))
	for i := 0; i < len(chars); i++ {
		if s[i] != '?' {
			chars[i] = s[i]
			continue
		}
		for j := byte('a'); j <= 'z'; j++ {
			if (i == 0 || j != chars[i-1]) && (i == len(chars)-1 || j != s[i+1]) {
				chars[i] = j
				break
			}
		}
	}
	return string(chars)
}
