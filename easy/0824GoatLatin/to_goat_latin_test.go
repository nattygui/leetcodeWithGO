package goatlatin

import "testing"

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				S: "I speak Goat Latin",
			},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.S); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
