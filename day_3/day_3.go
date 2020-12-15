package day_3

import (
	"fmt"
	"github.com/bojeanson/advent_of_code/utils"
)

func ResolveDay3(filename string) int {
	slopes := []slope{{right: 1, down: 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	encounteredTrees := 1
	for _, inputSlope := range slopes {
		encounteredTrees *= checkSlope(filename, inputSlope)
	}
	return encounteredTrees
}

func checkSlope(filename string, inputSlope slope) int {
	posX, posY := 0, 0
	tmpPosX := posX
	countTree := 0
	currentRow := ""
	for inputString := range utils.StringIteratorFromFile(filename) {
		if posY == 0 || posY%inputSlope.down != 0 {
			posY += 1
			fmt.Println(inputString)
			continue
		}
		posX += inputSlope.right
		tmpPosX = posX
		if posX >= len(inputString) {
			tmpPosX = posX % len(inputString)
		}
		switch inputString[tmpPosX] {
		case '#':
			countTree += 1
			currentRow = utils.ReplaceAtIndex(inputString, 'X', tmpPosX)
			fmt.Println(currentRow)
		case '.':
			currentRow = utils.ReplaceAtIndex(inputString, 'O', tmpPosX)
			fmt.Println(currentRow)
		}
		posY += 1
	}
	return countTree
}

type slope struct {
	right int
	down  int
}
