package array

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{2, 1, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
