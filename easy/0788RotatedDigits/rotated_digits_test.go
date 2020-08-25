package rotateddigits

import "testing"

func Test_rotatedDigits2(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{N: 5555},
			want: 1147,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotatedDigits2(tt.args.N); got != tt.want {
				t.Errorf("rotatedDigits2() = %v, want %v", got, tt.want)
			}
		})
	}
}
