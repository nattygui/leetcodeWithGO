package dynamicprogramming

import "testing"

func Test_countSubstrings(t *testing.T) {
	type question struct {
		name string
		arg  string
		want int
	}

	var questions = []question{
		{
			name: "test1",
			arg:  "abc",
			want: 3,
		},
		{
			name: "test2",
			arg:  "aaa",
			want: 6,
		},
		{
			name: "test3",
			arg:  "",
			want: 0,
		},
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := countSubstrings(q.arg); got != q.want {
				t.Errorf("countSubstrings() got: %d, but got: %d", q.want, got)
			}
		})
	}
	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := countSubstrings1(q.arg); got != q.want {
				t.Errorf("countSubstrings1() want: %d, but got: %d", q.want, got)
			}
		})
	}
}
