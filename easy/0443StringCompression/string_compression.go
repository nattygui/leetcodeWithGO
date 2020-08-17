package _443StringCompression

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}


	i, j, temp, count := 0, 0, chars[0], 0

	for i < len(chars) {
		if chars[i] == temp {
			i++
			count++
		} else if chars[i] != temp && count != 1 {
			chars[j] = temp
			j++
			if count > 9 {
				tempChars := make([]byte, 0)
				for count != 0 {
					tempChars = append(tempChars, byte((count % 10) + '0'))
					count /= 10
				}
				for k := len(tempChars)-1; k >= 0; k-- {
					chars[j] = tempChars[k]
					j++
				}
			} else {
				chars[j] = byte(count + '0')
				j++
			}
			temp = chars[i]
			count = 0
		} else {
			chars[j] = temp
			j++
			temp = chars[i]
			count = 0
		}
	}


	if count != 1 {
		chars[j] = temp
		j++
		if count > 9 {
			tempChars := make([]byte, 0)
			for count != 0 {
				tempChars = append(tempChars, byte((count % 10) + '0'))
				count /= 10
			}
			for k := len(tempChars)-1; k >= 0; k-- {
				chars[j] = tempChars[k]
				j++
			}
		} else {
			chars[j] = byte(count + '0')
			j++
		}
	} else {
		chars[j] = temp
		j++
	}

	return j
}
