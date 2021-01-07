package dynamicprogramming

import "testing"

func Test_maxProfit5(t *testing.T) {
	type question struct {
		name string
		arg1 []int
		arg2 int
		want int
	}
	questions := []question{
		{
			name: "test1",
			arg1: []int{1, 3, 2, 8, 4, 9},
			arg2: 2,
			want: 8,
		},
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := maxProfit5(q.arg1, q.arg2); got != q.want {
				t.Errorf("maxProfit5() want: %d, but got: %d", q.want, got)
			}
		})
	}
}
