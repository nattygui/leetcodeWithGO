package dynamicprogramming

import "testing"

func Test_maxProfit1(t *testing.T) {
	type question struct {
		name string
		arg  []int
		want int
	}
	questions := []question{
		{
			name: "test1",
			arg:  []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			name: "test2",
			arg:  []int{7, 6, 4, 3, 1},
			want: 0,
		},
		{
			name: "test3",
			arg:  []int{},
			want: 0,
		},
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := maxProfit1(q.arg); got != q.want {
				t.Errorf("maxProfit1() want: %d, but got: %d", q.want, got)
			}

			if got := maxProfit2(q.arg); got != q.want {
				t.Errorf("maxProfit2() want: %d, but got: %d", q.want, got)
			}
		})
	}
}
