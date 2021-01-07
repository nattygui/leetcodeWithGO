package dynamicprogramming

import "testing"

func Test_maxProfit3(t *testing.T) {
	type question struct {
		name string
		arg  []int
		want int
	}
	questions := []question{
		{
			name: "test1",
			arg:  []int{7, 1, 5, 3, 6, 4},
			want: 7,
		},
		{
			name: "test2",
			arg:  []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			name: "test3",
			arg:  []int{},
			want: 0,
		},
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := maxProfit3(q.arg); got != q.want {
				t.Errorf("maxProfit3() want: %d, but got: %d", q.want, got)
			}
		})
	}
}
