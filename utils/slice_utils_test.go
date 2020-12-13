package utils

import (
	"reflect"
	"testing"
)

func TestFindIntPosInSlice(t *testing.T) {
	type args struct {
		searchedValue int
		list          []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Should find searched int at position 2", args: struct {
			searchedValue int
			list          []int
		}{searchedValue: 5, list: []int{1, 2, 5, 4}}, want: 2, wantErr: false},
		{name: "Should not find searched int because not present", args: struct {
			searchedValue int
			list          []int
		}{searchedValue: 5, list: []int{1, 2, 3, 4}}, want: 0, wantErr: true},
		{name: "Should not find searched int because empty  slice", args: struct {
			searchedValue int
			list          []int
		}{searchedValue: 5, list: []int{}}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindIntPosInSlice(tt.args.searchedValue, tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindIntPosInSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindIntPosInSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertInSliceAtPosition(t *testing.T) {
	type args struct {
		lines    []int
		i        int
		inputInt int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Should insert int at position 4", args: struct {
			lines    []int
			i        int
			inputInt int
		}{lines: []int{0, 1, 2, 3, 5}, i: 4, inputInt: 4}, want: []int{0, 1, 2, 3, 4, 5}},
		{name: "Should not insert int at negative position", args: struct {
			lines    []int
			i        int
			inputInt int
		}{lines: []int{0, 1, 2, 3, 4}, i: -1, inputInt: 5}, want: []int{0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInSliceAtPosition(tt.args.lines, tt.args.i, tt.args.inputInt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertInSliceAtPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopFirstElementOfSlice(t *testing.T) {
	type args struct {
		lines []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   []int
		wantErr bool
	}{
		{name: "Should return 4 and smaller slice", args: struct{ lines []int }{lines: []int{4, 0, 1, 2, 3}}, want: 4,
			want1: []int{0, 1, 2, 3}, wantErr: false},
		{name: "Should return error message with empty slice", args: struct{ lines []int }{lines: []int{}}, want: 0,
			want1: []int{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PopFirstElementOfSlice(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopFirstElementOfSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PopFirstElementOfSlice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopFirstElementOfSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
