package thousandseparator

func thousandSeparator(n int) string {
	results := make([]byte, 0)
	num := 0
	for {
		results = append(results, byte(n%10)+'0')
		num++
		n = n / 10
		if n == 0 {
			break
		}
		if num%3 == 0 {
			results = append(results, '.')
		}
	}
	length := len(results)
	for i := 0; i < length/2; i++ {
		results[i], results[length-i-1] = results[length-i-1], results[i]
	}
	return string(results)
}
