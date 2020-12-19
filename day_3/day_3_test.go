package day_3

import "testing"

func Test_checkSlope(t *testing.T) {
	type args struct {
		filename   string
		inputSlope slope
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should encounter 1 tree with right 1 and down 1",
			args: args{
				filename:   "test_input_part_1.txt.txt",
				inputSlope: slope{1, 1},
			},
			want: 1,
		},
		{
			name: "Should encounter 4 tree with right 3 and down 1",
			args: args{
				filename:   "test_input_part_1.txt.txt",
				inputSlope: slope{3, 1},
			},
			want: 4,
		},
		{
			name: "Should encounter 0 tree with right 1 and down 2",
			args: args{
				filename:   "test_input_part_1.txt.txt",
				inputSlope: slope{1, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSlope(tt.args.filename, tt.args.inputSlope); got != tt.want {
				t.Errorf("checkSlope() = %v, want %v", got, tt.want)
			}
		})
	}
}
