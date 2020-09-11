package mds

func minDeletionSize(A []string) int {
	ySize := len(A)
	xSize := len(A[0])
	res := make([]int, 0)

	for i := 0; i < xSize; i++ {
		for j := 1; j < ySize; j++ {
			if A[j][i] < A[j-1][i] {
				res = append(res, i)
				break
			}
		}
	}
	return len(res)
}
