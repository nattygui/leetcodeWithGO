package _125ValidPalindrome

func isPalindrome(s string) bool {
	newstr := clearNonNumberAndLetter(s)
	length := len(newstr)
	for i := 0; i < length/2; i++{
		if newstr[i] != newstr[length-i-1] {
			return false
		}
	}
	return true
}

func clearNonNumberAndLetter(str string) string {
	newByte := []byte(str)
	for i := 0; i < len(newByte); i++ {
		if ('a' <= newByte[i] && newByte[i] <= 'z') || ('0' <= newByte[i] && newByte[i] <= '9') {
			continue
		} else if 'A' <= newByte[i] && newByte[i] <= 'Z' {
			newByte[i] = newByte[i] + 32
		} else {
			newByte = append(newByte[0:i], newByte[i+1:]...)
			i--
		}
	}
	return string(newByte)
}