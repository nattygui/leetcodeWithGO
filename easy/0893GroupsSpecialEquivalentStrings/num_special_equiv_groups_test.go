package numspecialequivgroups

import "testing"

func Test_numSpecialEquivGroups(t *testing.T) {
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
				A: []string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSpecialEquivGroups(tt.args.A); got != tt.want {
				t.Errorf("numSpecialEquivGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
