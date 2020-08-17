package _415AddStrings

func addStrings(num1 string, num2 string) string {
	res := make([]byte, 0)

	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0

	t1 := 0
	t2 := 0

	for i >= 0 || j >= 0 {
		if i >= 0 {
			t1 = int(num1[i] - '0')
		} else {
			t1 = 0
		}
		if j >= 0 {
			t2 = int(num2[j] - '0')
		} else {
			t2 = 0
		}
		tmp := t1 + t2 + carry
		carry = tmp / 10
		res = append(res, byte(tmp % 10 + '0'))
		i--
		j--
	}

	if carry != 0 {
		res = append(res, byte(carry + '0'))
	}

	for k := 0; k < len(res) / 2; k++ {
		temp := res[k]
		res[k] = res[len(res) - k - 1]
		res[len(res) - k - 1] = temp
	}

	return string(res)

}