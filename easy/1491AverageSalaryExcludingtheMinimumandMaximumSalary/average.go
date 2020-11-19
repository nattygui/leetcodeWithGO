package average

func average(salary []int) float64 {
	sum := salary[0]
	min := salary[0]
	max := salary[0]
	length := len(salary)
	for i := 1; i < length; i++ {
		sum += salary[i]
		if min > salary[i] {
			min = salary[i]
		}
		if max < salary[i] {
			max = salary[i]
		}
	}
	return float64(sum-min-max) / float64(length-2)
}
