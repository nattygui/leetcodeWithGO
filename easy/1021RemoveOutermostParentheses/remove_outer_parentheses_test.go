package removeouterparentheses

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				S: "(()())(())",
			},
			want: "()()()",
		},
		{
			name: "Test2",
			args: args{
				S: "(()())(())(()(()))",
			},
			want: "()()()()(())",
		},
		{
			name: "Test3",
			args: args{
				S: "()()",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.S); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
