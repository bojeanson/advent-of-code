package day_4

import (
	"reflect"
	"testing"
)

func TestResolveDay4(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should",
			args: args{
				filename: "test_input.txt",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveDay4(tt.args.filename); got != tt.want {
				t.Errorf("ResolveDay4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPassportValidity(t *testing.T) {
	type args struct {
		currentPassport passport
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPassportValidity(tt.args.currentPassport); got != tt.want {
				t.Errorf("checkPassportValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validPid(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Should be a valid pid", args: args{inputString: "pid:000000001"}, want: true},
		{name: "Should be a invalid pid", args: args{inputString: "pid:0123456789"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPid(tt.args.inputString); got != tt.want {
				t.Errorf("validPid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateByr(t *testing.T) {
	type args struct {
		inputString     string
		currentPassport passport
	}
	tests := []struct {
		name string
		args args
		want passport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateByr(tt.args.inputString, tt.args.currentPassport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateByr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateEcl(t *testing.T) {
	type args struct {
		inputString     string
		currentPassport passport
	}
	tests := []struct {
		name string
		args args
		want passport
	}{
		{name: "Should be a valid ecl", args: args{inputString: "ecl:brn", currentPassport: passport{}}, want: passport{ecl: true}},
		{name: "Should be a invalid ecl", args: args{inputString: "ecl:wat", currentPassport: passport{}}, want: passport{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateEcl(tt.args.inputString, tt.args.currentPassport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateEcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateEyr(t *testing.T) {
	type args struct {
		inputString     string
		currentPassport passport
	}
	tests := []struct {
		name string
		args args
		want passport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateEyr(tt.args.inputString, tt.args.currentPassport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateEyr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHcl(t *testing.T) {
	type args struct {
		inputString     string
		currentPassport passport
	}
	tests := []struct {
		name string
		args args
		want passport
	}{
		{name: "Should be a valid hcl", args: args{inputString: "hcl:#123abc", currentPassport: passport{}}, want: passport{hcl: true}},
		{name: "Should be a invalid hcl", args: args{inputString: "hcl:#123abz", currentPassport: passport{}}, want: passport{}},
		{name: "Should also be a invalid hcl", args: args{inputString: "hcl:123abc", currentPassport: passport{}}, want: passport{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHcl(tt.args.inputString, tt.args.currentPassport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateHcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHgt(t *testing.T) {
	type args struct {
		inputString     string
		currentPassport passport
	}
	tests := []struct {
		name string
		args args
		want passport
	}{
		{name: "Should be a valid hgt", args: args{inputString: "hgt:60in", currentPassport: passport{}}, want: passport{hgt: true}},
		{name: "Should also be a valid hgt", args: args{inputString: "hgt:190cm", currentPassport: passport{}}, want: passport{hgt: true}},
		{name: "Should be a invalid hgt", args: args{inputString: "hgt:190in", currentPassport: passport{}}, want: passport{hgt: false}},
		{name: "Should also be a invalid hgt", args: args{inputString: "hgt:190", currentPassport: passport{}}, want: passport{hgt: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHgt(tt.args.inputString, tt.args.currentPassport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateHgt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateIyr(t *testing.T) {
	type args struct {
		inputString     string
		currentPassport passport
	}
	tests := []struct {
		name string
		args args
		want passport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateIyr(tt.args.inputString, tt.args.currentPassport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateIyr() = %v, want %v", got, tt.want)
			}
		})
	}
}
