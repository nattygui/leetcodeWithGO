package dynamicprogramming

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type question struct {
		name string
		arg  string
		want string
	}

	var questions = []question{
		{
			name: "test1",
			arg:  "babad",
			want: "bab",
		},
		{
			name: "test2",
			arg:  "babad",
			want: "bab",
		},
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := longestPalindrome(q.arg); got != q.want {
				t.Errorf("want: %s, but got %s", q.want, got)
			}
		})
	}

	for _, q := range questions {
		t.Run(q.name, func(t *testing.T) {
			if got := longestPalindrome1(q.arg); got != q.want {
				t.Errorf("want: %s, but got %s", q.want, got)
			}
		})
	}
}
