package _168ExcelSheetColumnTitle

func convertToTitle(n int) string {
	if n <= 26 {
		return string('A'+n-1)
	}
	newByte := []byte{}
	var temp int
	for n != 0 {
		temp = n % 26
		n = n / 26
		newByte = append(newByte, byte('A'+temp-1))
	}

	var temp1 byte
	for i := 0; i < len(newByte) / 2; i++ {
		temp1 = newByte[i]
		newByte[i] = newByte[len(newByte)-i-1]
		newByte[len(newByte)-i-1] = temp1
	}

	return string(newByte)
}