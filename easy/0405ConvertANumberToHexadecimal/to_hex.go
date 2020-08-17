package _405ConvertANumberToHexadecimal

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += 1 << 32
	}

	bytes := make([]byte, 0)


	for num != 0 {
		temp := num % 16
		if temp <= 10 {
			bytes = append(bytes, byte(temp+'0'))
		} else {
			bytes = append(bytes, byte(temp-10+'a'))
		}
		num /= 16
	}

	for i := 0; i < len(bytes)/2; i++ {
		temp := bytes[i]
		bytes[i] = bytes[len(bytes)-i-1]
		bytes[len(bytes)-i-1] = temp
	}

	return string(bytes)
}
