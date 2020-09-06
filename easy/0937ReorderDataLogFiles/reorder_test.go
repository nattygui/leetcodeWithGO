package reorder

import (
	"reflect"
	"testing"
)

func Test_reorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				logs: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			},
			want: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderLogFiles(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
