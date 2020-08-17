package _344ReverseString

func reverseString(s []byte)  {
	length := len(s)

	var temp byte
	for i := 0; i < length / 2; i++ {
		temp = s[i]
		s[i] = s[length - i - 1]
		s[length - i - 1] = temp
	}
}
