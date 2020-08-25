package toeplitzmatrix

import "testing"

func TestOK(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}

	res := isToeplitzMatrix(matrix)

	if !res {
		t.Error("error")
	}
}
