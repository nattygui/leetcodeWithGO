package maxdist

func maxDistToClosest(seats []int) int {
	indexs := make([]int, 0)

	for index, seat := range seats {
		if seat == 1 {
			indexs = append(indexs, index)
		}
	}

	lens := len(seats)
	res := 0
	for i, index := range indexs {
		if i == 0 {
			if res < index {
				res = index
			}
		}

		if i == len(indexs)-1 {
			if res < lens-index {
				res = lens - index
			}
		}

		if i != 0 {
			if res < (index-indexs[i-1])/2 {
				res = (index - indexs[i-1]) / 2
			}
		}

	}
	return res
}
