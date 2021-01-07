package dynamicprogramming

import "testing"

func Test_maxProfit4(t *testing.T) {
	type question struct {
		name string
		arg  []int
		want int
	}
	questions := []question{
		{
			name: "test1",
			arg:  []int{1, 2, 3, 0, 2},
			want: 3,
		},
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := maxProfit4(q.arg); got != q.want {
				t.Errorf("maxProfit4() want: %d, but got: %d", q.want, got)
			}
		})
	}
}
