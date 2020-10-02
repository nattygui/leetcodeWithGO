package mincosttomovechips

func minCostToMoveChips(position []int) int {
	if len(position) == 0 {
		return 0
	}

	oddNums := 0
	evenNums := 0
	for _, num := range position {
		if num%2 == 0 {
			evenNums++
			continue
		}
		oddNums++
	}

	if evenNums > oddNums {
		return evenNums
	}
	return oddNums
}
