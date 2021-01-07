package calcequation

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				equations: [][]string{{"a", "c"}, {"b", "e"}, {"c", "d"}, {"e", "d"}},
				values:    []float64{2.0, 3.0, 0.5, 5.0},
				queries:   [][]string{{"a", "b"}},
			},
			want: []float64{1.0 / 15.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %0.5f, want %0.5f", got, tt.want)
			}
		})
	}
}
