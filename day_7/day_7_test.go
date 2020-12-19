package day_7

import (
	"testing"
)

func TestResolveDay7(t *testing.T) {
	type args struct {
		part     string
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Should return 4 bag colors", args: args{part: "1", filename: "test_input_part_1.txt"}, want: 4},
		{name: "Should return 126 bag colors", args: args{part: "2", filename: "test_input_part_2.txt"}, want: 126},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveDay7(tt.args.part, tt.args.filename); got != tt.want {
				t.Errorf("ResolveDay7() = %v, want %v", got, tt.want)
			}
		})
	}
}
