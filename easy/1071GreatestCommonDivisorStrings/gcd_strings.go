package gcd

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	remainder := modOfStrings(str1, str2)
	if remainder == "" {
		return str2
	} else if remainder == str1 {
		return ""
	}
	return gcdOfStrings(str2, remainder)
}

func modOfStrings(str1 string, str2 string) string {
	length := len(str2)
	remainder := str1
	for {
		if len(str1) < length || remainder[0:length] != str2 {
			break
		}
		remainder = remainder[length:]
	}
	return remainder
}
