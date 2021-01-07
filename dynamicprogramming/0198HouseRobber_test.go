package dynamicprogramming

import "testing"

func Test_rob(t *testing.T) {
	type question struct {
		name string
		arg  []int
		want int
	}
	questions := []question{
		{
			name: "test1",
			arg:  []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "test2",
			arg:  []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "test3",
			arg:  []int{},
			want: 0,
		},
		{
			name: "test4",
			arg:  []int{2},
			want: 2,
		},
		{
			name: "test5",
			arg:  []int{2, 3},
			want: 3,
		},
	}
	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := rob(q.arg); got != q.want {
				t.Errorf("rob() want: %d, but got: %d", q.want, got)
			}
		})
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := rob1(q.arg); got != q.want {
				t.Errorf("rob() want: %d, but got: %d", q.want, got)
			}
		})
	}
}
