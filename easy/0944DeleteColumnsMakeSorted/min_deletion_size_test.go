package mds

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				A: []string{"cda", "daf", "ghi"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.A); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
