package findpivotindex

import "testing"

func TestOK(t *testing.T) {
	nums := [][]int{
		{1, 7, 3, 6, 5, 6},
		{-1, -1, -1, 0, 1, 1},
		{1, 2, 3},
	}

	results := []int{3, 0, -1}

	for i, numList := range nums {
		res := pivotIndex(numList)
		if res != results[i] {
			t.Errorf("res should be %v, not be %v", results[i], res)
		}
	}

}
