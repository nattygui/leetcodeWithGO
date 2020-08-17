package selfdividingnumbers

import "testing"

func TestOK(t *testing.T) {
	left := 1
	right := 22

	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}

	res := selfDividingNumbers(left, right)

	if len(res) != len(result) {
		t.Errorf("error1 res: %#v", res)
		return
	}

	for i := 0; i < len(res); i++ {
		if res[i] != result[i] {
			t.Errorf("error2 res: %#v", res)
		}
	}
}
