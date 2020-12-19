package day_8

import (
	"github.com/bojeanson/advent_of_code/utils"
	"regexp"
	"strconv"
	"strings"
)

func ResolveDay8(part string, filename string) int {
	var instructions []instruction
	for inputString := range utils.StringIteratorFromFile(filename) {
		instructions = append(instructions, parseInputLine(inputString))
	}
	return infiniteLoop(instructions)
}

func infiniteLoop(instructions []instruction) int {
	index := 0
	var acc int
	visitedIndex := map[int]int{}
	ok := false
	for {
		switch instructions[index].operator {
		case "nop":
			visitedIndex, ok = updateVisitedIndex(visitedIndex, index)
			if ok {
				return acc
			}
			index += 1
		case "acc":
			visitedIndex, ok = updateVisitedIndex(visitedIndex, index)
			if ok {
				return acc
			}
			if instructions[index].sign == "+" {
				acc += instructions[index].value
			} else {
				acc -= instructions[index].value
			}
			index += 1
		case "jmp":
			visitedIndex, ok = updateVisitedIndex(visitedIndex, index)
			if ok {
				return acc
			}
			if instructions[index].sign == "+" {
				index += instructions[index].value
			} else {
				index -= instructions[index].value
			}
		}
	}
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
