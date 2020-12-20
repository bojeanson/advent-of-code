package day_9

import "testing"

func TestResolveDay9(t *testing.T) {
	type args struct {
		part     string
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Should return 127", args: args{part: "1", filename: "test_input.txt"}, want: 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveDay9(tt.args.part, tt.args.filename); got != tt.want {
				t.Errorf("ResolveDay9() = %v, want %v", got, tt.want)
			}
		})
	}
}
