package islongpressedname

func isLongPressedName(name string, typed string) bool {
	i := 0
	index := 0
	for i < len(name) {
		if index >= len(typed) {
			return false
		}
		if name[i] != typed[index] {
			return false
		}
		num1 := 1
		for i+1 < len(name) && name[i] == name[i+1] {
			i++
			num1++
		}
		num2 := 1
		for index+1 < len(typed) && typed[index] == typed[index+1] {
			index++
			num2++
		}

		if num1 > num2 {
			return false
		}

		i++
		index++
	}
	return true
}
