package findspecialinteger

func findSpecialInteger(arr []int) int {
	arrayLen := len(arr)
    if arrayLen == 1 {
        return arr[0]
    }    
	quarterLen := arrayLen / 4

	preNum := arr[0]
	tempLen := 1
	for i := 1; i < arrayLen; i++ {
		if arr[i] == preNum {
			tempLen++
			if tempLen > quarterLen {
				return arr[i]
			}
			continue
		} 
		preNum = arr[i]
		tempLen = 1
	}
	return -1
}