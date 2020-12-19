package day_7

import "testing"

func TestResolveDay7(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Should return 4 bag colors", args: args{filename: "test_input.txt"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveDay7(tt.args.filename); got != tt.want {
				t.Errorf("ResolveDay7() = %v, want %v", got, tt.want)
			}
		})
	}
}
