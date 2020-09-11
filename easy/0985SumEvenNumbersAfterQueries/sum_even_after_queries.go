package sumevenafterqueries

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, 0)
	evenSum := getEvenSum(A)
	temp := 0
	for _, query := range queries {
		index := query[1]
		value := query[0]

		if A[index]%2 == 0 {
			evenSum -= A[index]
		}

		temp = A[index] + value
		if temp%2 == 0 {
			evenSum += temp
		}
		A[index] = temp

		res = append(res, evenSum)
	}
	return res
}

func getEvenSum(A []int) int {
	sum := 0
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}

	return sum
}
