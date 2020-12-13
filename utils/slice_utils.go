package utils

import (
	"errors"
)

func PopFirstElementOfSlice(lines []int) (int, []int, error) {
	if len(lines) == 0 {
		err := errors.New("Empty slice provided")
		return 0, lines, err
	}
	var value int
	value, lines = lines[0], lines[1:]
	return value, lines, nil
}

func InsertInSliceAtPosition(lines []int, i int, inputInt int) []int {
	if i == -1 {
		lines = append(lines, inputInt)
	} else {
		lines = append(lines, 0)
		copy(lines[i+1:], lines[i:])
		lines[i] = inputInt
	}
	return lines
}

func FindIntPosInSlice(searchedValue int, list []int) (int, error) {
	for i, element := range list {
		if element == searchedValue {
			return i, nil
		}
	}
	return -1, errors.New("searched int not present in slice")
}
