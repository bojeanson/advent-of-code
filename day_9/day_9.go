package day_9

import (
	"github.com/bojeanson/advent_of_code/utils"
)

func ResolveDay9(part string, filename string) int {
	var numbers []int
	var preamble []int
	index := 0
	preambleLength := 25
	if filename == "test_input.txt" {
		preambleLength = 5
	}
	for inputInt := range utils.IntIteratorFromFile(filename) {
		numbers = append(numbers, inputInt)
		if len(preamble) < preambleLength {
			preamble = append(preamble, inputInt)
		} else {
			if !isValidNumber(preamble, inputInt) {
				if part == "1" {
					return inputInt
				} else {
					contiguousRange := computeContiguousRange(numbers[:index], 0, 1, inputInt)
					return min(contiguousRange) + max(contiguousRange)
				}
			}
			_, preamble, _ = utils.PopFirstElementOfSlice(preamble)
			preamble = append(preamble, inputInt)
		}
		index += 1
	}
	return -1
}

func computeContiguousRange(numbers []int, i int, j int, invalidNumber int) []int {
	subNumbers := numbers[i:j]
	sumValue := sum(subNumbers)
	if sumValue < invalidNumber {
		subNumbers = computeContiguousRange(numbers, i, j+1, invalidNumber)
	}
	if sumValue > invalidNumber {
		subNumbers = computeContiguousRange(numbers, i+1, j, invalidNumber)
	}
	if sumValue == invalidNumber {
		return subNumbers
	}
	return subNumbers
}

func isValidNumber(preamble []int, inputInt int) bool {
	var value int
	for len(preamble) > 0 {
		value, preamble, _ = utils.PopFirstElementOfSlice(preamble)
		searchedValue := inputInt - value
		pos, _ := utils.FindIntPosInSlice(searchedValue, preamble)
		if pos >= 0 {
			return true
		}
	}
	return false
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
func min(array []int) int {
	minVal := array[0]
	for _, v := range array {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}
func max(array []int) int {
	maxVal := array[0]
	for _, v := range array {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
