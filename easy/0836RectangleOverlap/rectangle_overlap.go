package rectangleoverlap

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 如果rec2在rec1右边
	if rec2[0] >= rec1[2] {
		return false
	} else if rec2[2] <= rec1[0] { // 左边
		return false
	} else if rec2[1] >= rec1[3] { // 上边
		return false
	} else if rec2[3] <= rec1[1] {
		return false
	}

	return true
}
