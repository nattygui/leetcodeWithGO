package replaceelements

func replaceElements(arr []int) []int {
	length := len(arr) - 1
	maxNum := arr[length]
	arr[length] = -1
	for i := length - 1; i >= 0; i-- {
		tempNum := arr[i]
		arr[i] = maxNum
		maxNum = findMaxNum(maxNum, tempNum)
	}

	return arr
}

func findMaxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
