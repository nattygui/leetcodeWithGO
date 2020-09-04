package numberoflines

import (
	"reflect"
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	type args struct {
		widths []int
		S      string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				widths: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				S:      "abcdefghijklmnopqrstuvwxyz",
			},
			want: []int{3, 60},
		},
		{
			name: "Test2",
			args: args{
				widths: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				S:      "bbbcccdddaaa",
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfLines(tt.args.widths, tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
