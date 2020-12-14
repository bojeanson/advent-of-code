package day_2

import "testing"

func Test_countValidPassword(t *testing.T) {
	type args struct {
		lines []string
		part  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Should return 2", args: args{part: "1", lines: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}}, want: 2},
		{name: "Should return 1", args: args{part: "2", lines: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidPassword(tt.args.part, tt.args.lines); got != tt.want {
				t.Errorf("countValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Should return true with valid password1", args: args{password: "1-3 a: abcde"}, want: true},
		{name: "Should return false with invalid password2", args: args{password: "1-3 b: cdefg"}, want: false},
		{name: "Should return true with valid password3", args: args{password: "2-9 c: ccccccccc"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPassword(tt.args.password, nil); got != tt.want {
				t.Errorf("isValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
