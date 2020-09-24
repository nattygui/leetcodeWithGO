package prefixesdivby5

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				A: []int{0, 1, 1},
			},
			want: []bool{true, false, false},
		},
		{
			name: "test2",
			args: args{
				A: []int{1, 1, 1},
			},
			want: []bool{false, false, false},
		},
		{
			name: "test3",
			args: args{
				A: []int{0, 1, 1, 1, 1, 1},
			},
			want: []bool{true, false, false, false, true, false},
		},
		{
			name: "test4",
			args: args{
				A: []int{1, 1, 1, 0, 1},
			},
			want: []bool{false, false, false, false, false},
		},
		{
			name: "test5",
			args: args{
				A: []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0},
			},
			want: []bool{false, false, true, false, false, false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, true, false, false, true, true, true, true, true, true, true, false, false, true, false, false, false, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixesDivBy5(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
