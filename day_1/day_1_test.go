package day_1

import (
	"reflect"
	"testing"
)

func TestFindSucceedingPair1(t *testing.T) {
	type args struct {
		lines []int
	}
	tests := []struct {
		name string
		args args
		want Pair
	}{
		{name: "Should", args: struct{ lines []int }{lines: []int{0, 2020}}, want: Pair{0, 2020}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSucceedingPair(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSucceedingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSucceedingTriplet(t *testing.T) {
	type args struct {
		lines []int
	}
	tests := []struct {
		name string
		args args
		want Triplet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSucceedingTriplet(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSucceedingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveDay1(t *testing.T) {
	type args struct {
		inputFilePath string
	}
	tests := []struct {
		name  string
		args  args
		want  Pair
		want1 Triplet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ResolveDay1(tt.args.inputFilePath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveDay1() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ResolveDay1() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
