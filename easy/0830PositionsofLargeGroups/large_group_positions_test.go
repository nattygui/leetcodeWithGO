package largegrouppositions

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				S: "abcdddeeeeaabbb",
			},
			want: [][]int{
				{3, 5},
				{6, 9},
				{12, 14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
