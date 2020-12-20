package day_8

import (
	"github.com/bojeanson/advent_of_code/utils"
	"regexp"
	"strconv"
	"strings"
)

func ResolveDay8(part string, filename string) int {
	var originalInstructions []instruction
	for inputString := range utils.StringIteratorFromFile(filename) {
		originalInstructions = append(originalInstructions, parseInputLine(inputString))
	}
	runInstructionIndexes, _, acc := infiniteLoop(originalInstructions)
	if part == "1" {
		return acc
	}
	acc = fixProgram(originalInstructions, runInstructionIndexes)
	return acc
}

func infiniteLoop(instructions []instruction) ([]int, bool, int) {
	index := 0
	var acc int
	visitedIndex := map[int]int{}
	ok := false
	var runInstructionIndexes []int
	for index < len(instructions) {
		visitedIndex, ok = updateVisitedIndex(visitedIndex, index)
		if ok {
			return runInstructionIndexes, true, acc
		}
		switch instructions[index].operator {
		case "nop":
			runInstructionIndexes = append(runInstructionIndexes, index)
			index += 1
		case "acc":
			if instructions[index].sign == "+" {
				acc += instructions[index].value
			} else {
				acc -= instructions[index].value
			}
			runInstructionIndexes = append(runInstructionIndexes, index)
			index += 1
		case "jmp":
			runInstructionIndexes = append(runInstructionIndexes, index)
			if instructions[index].sign == "+" {
				index += instructions[index].value
			} else {
				index -= instructions[index].value
			}
		}
	}
	return runInstructionIndexes, false, acc
}

func fixProgram(originalInstructions []instruction, runInstructionIndexes []int) int {
	wasInfiniteLoop := true
	acc := 0
	for wasInfiniteLoop {
		for i := 0; i < len(runInstructionIndexes); i++ {
			acc = 0
			instructionIndex := runInstructionIndexes[i]
			fixedInstructions := make([]instruction, len(originalInstructions))
			copy(fixedInstructions, originalInstructions)
			switch fixedInstructions[instructionIndex].operator {
			case "jmp":
				fixedInstructions[instructionIndex].operator = "nop"
			case "nop":
				fixedInstructions[instructionIndex].operator = "jmp"
			default:
				continue
			}
			_, wasInfiniteLoop, acc = infiniteLoop(fixedInstructions)
			if !wasInfiniteLoop {
				break
			}
		}
	}
	return acc
}

func updateVisitedIndex(visitedIndex map[int]int, index int) (map[int]int, bool) {
	_, ok := visitedIndex[index]
	if !ok {
		visitedIndex[index] = 1
	}
	return visitedIndex, ok
}

func parseInputLine(inputLine string) instruction {
	var instruc = instruction{}
	switch {
	case strings.Contains(inputLine, "acc"):
		instruc.operator = "acc"
		instruc = parseInstruction(inputLine, instruc)
	case strings.Contains(inputLine, "jmp"):
		instruc.operator = "jmp"
		instruc = parseInstruction(inputLine, instruc)
	case strings.Contains(inputLine, "nop"):
		instruc.operator = "nop"
		instruc = parseInstruction(inputLine, instruc)
	}
	return instruc
}

func parseInstruction(inputLine string, instruc instruction) instruction {
	re := regexp.MustCompile("(\\+|-)(\\d+)")
	matched := re.FindAllSubmatch([]byte(inputLine[3:]), -1)
	for _, match := range matched {
		instruc.sign = string(match[1])
		intAsInt64, _ := strconv.ParseInt(string(match[2]), 10, 64)
		instruc.value = int(intAsInt64)
	}
	return instruc
}

type instruction struct {
	operator string
	sign     string
	value    int
}
