package day_1

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"strings"
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
		{name: "Should return first and second element of array", args: struct{ lines []int }{lines: []int{0, 2020}},
			want: Pair{0, 2020}},
		{name: "Should return second and penultimate element of array",
			args: struct{ lines []int }{lines: []int{0, 1, 489, 2019, 2039}},
			want: Pair{1, 2019}},
		{name: "Should log that no Pair were found",
			args: struct{ lines []int }{lines: []int{0, 1, 489, 2018, 2039}},
			want: Pair{}},
	}
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSucceedingPair(tt.args.lines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSucceedingPair() = %v, want %v", got, tt.want)
			}
			if reflect.DeepEqual(got, Pair{}) && !strings.Contains(buf.String(), "No Pair found") {
				t.Errorf("findSucceedingPair() = %v, want %v", got, tt.want)
				t.Errorf("Expected log message= %v, want %v", buf.String(), "No Pair found")
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
		{name: "Should return first, second and third element of array", args: args{lines: []int{0, 1, 2019}}, want: Triplet{
			FirstMember:  0,
			SecondMember: 1,
			ThirdMember:  2019,
		}},
		{name: "Should return second, third and ultimate element of array", args: args{lines: []int{0, 9, 101, 290, 1910}}, want: Triplet{
			FirstMember:  9,
			SecondMember: 101,
			ThirdMember:  1910,
		}},
		{name: "Should log that no Triplet were found",
			args: args{lines: []int{0, 9, 105, 290, 1940}},
			want: Triplet{}},
	}
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSucceedingTriplet(tt.args.lines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSucceedingTriplet() = %v, want %v", got, tt.want)
			}
			if reflect.DeepEqual(got, Triplet{}) && !strings.Contains(buf.String(), "No Triplet found") {
				t.Errorf("findSucceedingTriplet() = %v, want %v", got, tt.want)
				t.Errorf("Expected log message= %v, want %v", buf.String(), "No Triplet found")
			}
		})
	}
}
