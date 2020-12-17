package day_5

import (
	"github.com/bojeanson/advent_of_code/utils"
)

func ResolveDay5(filename string) int {
	highestSeatId := 0
	plane := makePlane(8, 128)
	var seatIds []int
	for inputString := range utils.StringIteratorFromFile(filename) {
		seatPosition := position{frontRow: 0, backRow: 127, leftColumn: 0, rightColumn: 7}
		seatPosition = findSeatPosition(inputString, seatPosition)
		plane[seatPosition.frontRow][seatPosition.leftColumn] = "X"
		seatId := computeSeatId(seatPosition)
		seatIds = append(seatIds, seatId)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}
	mySeat := position{}
	for i, row := range plane {
		allElementsEqualToExceptOne, j := utils.AllElementsEqualToExceptOne(row, "X")
		if allElementsEqualToExceptOne {
			mySeat = position{
				frontRow:    i,
				backRow:     i,
				leftColumn:  j,
				rightColumn: j,
			}
		}
	}
	return computeSeatId(mySeat)
}

func makePlane(dx int, dy int) [][]string {
	array := make([][]string, dy)
	for i := range array {
		array[i] = []string{"O", "O", "O", "O", "O", "O", "O", "O"}
	}
	return array
}

func findSeatPosition(inputString string, seatPosition position) position {
	if len(inputString) == 0 {
		return seatPosition
	}
	switch inputString[0] {
	case 'F':
		seatPosition.backRow = (seatPosition.frontRow + seatPosition.backRow) / 2
		return findSeatPosition(inputString[1:], seatPosition)
	case 'B':
		seatPosition.frontRow = (seatPosition.frontRow+seatPosition.backRow)/2 + 1
		return findSeatPosition(inputString[1:], seatPosition)
	case 'L':
		seatPosition.rightColumn = (seatPosition.leftColumn + seatPosition.rightColumn) / 2
		return findSeatPosition(inputString[1:], seatPosition)
	case 'R':
		seatPosition.leftColumn = (seatPosition.leftColumn+seatPosition.rightColumn)/2 + 1
		return findSeatPosition(inputString[1:], seatPosition)
	}
	return seatPosition
}

func computeSeatId(seatPosition position) int {
	return seatPosition.frontRow*8 + seatPosition.leftColumn
}

type position struct {
	frontRow    int
	backRow     int
	leftColumn  int
	rightColumn int
}
