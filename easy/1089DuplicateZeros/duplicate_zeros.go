package duplicatezeros

func duplicateZeros(arr []int) {
	zeros := 0
	length := len(arr)
	for _, num := range arr {
		if num != 0 {
			continue
		}
		zeros++
	}

	for i := length - 1; i >= 0; i-- {
		if arr[i] == 0 {
			zeros--
		}

		if i+zeros < length {
			arr[i+zeros] = arr[i]
			if arr[i] == 0 {
				arr[i+zeros+1] = arr[i]
			}
		}
	}
}
