package day_5

import (
	"reflect"
	"testing"
)

func TestResolveDay5(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveDay5(tt.args.filename); got != tt.want {
				t.Errorf("ResolveDay5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSeatPosition(t *testing.T) {
	type args struct {
		inputString  string
		seatPosition position
	}
	tests := []struct {
		name string
		args args
		want position
	}{
		{name: "Should return row 44, column 5", args: args{
			inputString:  "FBFBBFFRLR",
			seatPosition: position{frontRow: 0, backRow: 127, leftColumn: 0, rightColumn: 7},
		}, want: position{frontRow: 44, backRow: 44, leftColumn: 5, rightColumn: 5}},
		{name: "Should return row 70, column 7", args: args{
			inputString:  "BFFFBBFRRR",
			seatPosition: position{frontRow: 0, backRow: 127, leftColumn: 0, rightColumn: 7},
		}, want: position{frontRow: 70, backRow: 70, leftColumn: 7, rightColumn: 7}},
		{name: "Should return row 14, column 7", args: args{
			inputString:  "FFFBBBFRRR",
			seatPosition: position{frontRow: 0, backRow: 127, leftColumn: 0, rightColumn: 7},
		}, want: position{frontRow: 14, backRow: 14, leftColumn: 7, rightColumn: 7}},
		{name: "Should return row 102, column 4", args: args{
			inputString:  "BBFFBBFRLL",
			seatPosition: position{frontRow: 0, backRow: 127, leftColumn: 0, rightColumn: 7},
		}, want: position{frontRow: 102, backRow: 102, leftColumn: 4, rightColumn: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSeatPosition(tt.args.inputString, tt.args.seatPosition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSeatPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeSeatId(t *testing.T) {
	type args struct {
		seatPosition position
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Should return 357", args: args{seatPosition: position{frontRow: 44, leftColumn: 5}}, want: 357},
		{name: "Should return 567", args: args{seatPosition: position{frontRow: 70, leftColumn: 7}}, want: 567},
		{name: "Should return 119", args: args{seatPosition: position{frontRow: 14, leftColumn: 7}}, want: 119},
		{name: "Should return 820", args: args{seatPosition: position{frontRow: 102, leftColumn: 4}}, want: 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeSeatId(tt.args.seatPosition); got != tt.want {
				t.Errorf("computeSeatId() = %v, want %v", got, tt.want)
			}
		})
	}
}
