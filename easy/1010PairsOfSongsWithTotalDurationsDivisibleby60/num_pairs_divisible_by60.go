package numpairsdivisibleby60

func numPairsDivisibleBy60(time []int) int {
	result := 0
	carrires := make([]int, 60)
	for _, value := range time {
		carry := value % 60
		carrires[carry]++
	}

	result += carrires[0] * (carrires[0] - 1) / 2
	result += carrires[30] * (carrires[30] - 1) / 2
	i := 1
	j := 59
	for i < j {
		result += carrires[i] * carrires[j]
		i++
		j--
	}

	return result
}
