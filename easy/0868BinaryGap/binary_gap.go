package binarygap

func binaryGap(N int) int {
	pre := -1

	index := 0
	res := 0
	for N != 0 {
		if N%2 == 1 {
			if pre == -1 {
				pre = index
			} else {
				if res < index-pre {
					res = index - pre
				}
				pre = index
			}
		}
		index++
		N = N >> 1
	}
	return res
}
