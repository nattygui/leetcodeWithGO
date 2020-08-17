package tolowercase

func toLowerCase(str string) string {
	bytes := []byte(str)
	for index, value := range bytes {
		if value >= 'A' && value <= 'Z' {
			bytes[index] += 32
		}
	}
	return string(bytes)
}
