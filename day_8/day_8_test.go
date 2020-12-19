package day_8

import "testing"

func TestResolveDay8(t *testing.T) {
	type args struct {
		part     string
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Should", args: args{part: "1", filename: "test_input_file.txt"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveDay8(tt.args.part, tt.args.filename); got != tt.want {
				t.Errorf("ResolveDay8() = %v, want %v", got, tt.want)
			}
		})
	}
}
