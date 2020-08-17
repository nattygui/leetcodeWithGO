package _541ReverseStringII

func reverseStr(s string, k int) string {
	res := []byte(s)

	for i := 0; i < len(res); i=i + 2 * k {
		end := 0
		if i + k > len(res) {
			end = len(res) - 1
		} else {
			end = i + k -1
		}
		for j := 0; j < (end - i + 1) / 2; j++ {
			res[i+j], res[end-j] = res[end-j], res[i+j]
		}
	}

	return string(res)
}
