package surfacearea

import "testing"

func Test_surfaceArea(t *testing.T) {
	type args struct {
		grid [][]int
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
				grid: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := surfaceArea(tt.args.grid); got != tt.want {
				t.Errorf("surfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
