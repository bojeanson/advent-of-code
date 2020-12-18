package day_6

import (
	"testing"
)

func TestResolveDay6(t *testing.T) {
	type args struct {
		part     string
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Should count 6 answers", args: args{part: "1", filename: "test_input.txt"}, want: 11},
		{name: "Should count 6 answers", args: args{part: "2", filename: "test_input.txt"}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveDay6(tt.args.part, tt.args.filename); got != tt.want {
				t.Errorf("ResolveDay6() = %v, want %v", got, tt.want)
			}
		})
	}
}
