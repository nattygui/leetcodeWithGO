package dynamicprogramming

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type question struct {
		name string
		arg  string
		want int
	}

	var questions = []question{
		{
			name: "test1",
			arg:  "",
			want: 0,
		},
		{
			name: "test2",
			arg:  "a",
			want: 1,
		},
		{
			name: "test3",
			arg:  "bbbab",
			want: 4,
		},
		{
			name: "test4",
			arg:  "cbbd",
			want: 2,
		},
	}
	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(q.arg); got != q.want {
				t.Errorf("longestPalindromeSubseq() got: %d, want: %d", got, q.want)
			}
		})
	}
}
