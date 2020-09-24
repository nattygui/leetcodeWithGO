package commonchars

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				A: []string{"cool", "lock", "cook"},
			},
			want: []string{"c", "o"},
		},
		{
			name: "test2",
			args: args{
				A: []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"},
			},
			want: []string{},
		},
		{
			name: "test3",
			args: args{
				A: []string{"bella", "label", "roller"},
			},
			want: []string{"e", "l", "l"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
