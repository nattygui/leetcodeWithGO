package reconstruct

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
