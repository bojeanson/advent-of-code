package day_6

import (
	"github.com/bojeanson/advent_of_code/utils"
)

func ResolveDay6(part string, filename string) int {
	ndAnswers := 0
	groupAnswers := map[string]int{}
	for inputString := range utils.StringIteratorFromFile(filename) {
		switch inputString {
		case "":
			ndAnswers += countGroupAnswers(groupAnswers, part)
			groupAnswers = map[string]int{}
		default:
			groupAnswers = processAnswers(inputString, groupAnswers, part)
		}
	}
	ndAnswers += countGroupAnswers(groupAnswers, part)
	return ndAnswers
}

func processAnswers(inputString string, answers map[string]int, part string) map[string]int {
	if part == "1" {
		return processAnswersAnyoneRule(inputString, answers)
	} else {
		return processAnswersEveryoneRule(inputString, answers)
	}

}

func countGroupAnswers(answers map[string]int, part string) int {
	if part == "1" {
		return countGroupAnswersAnyoneRule(answers)
	} else {
		return countGroupAnswersEveryoneRule(answers)
	}
}

func countGroupAnswersEveryoneRule(currentGroup map[string]int) int {
	nb_person := currentGroup["nb_person"]
	totalCount := 0
	for key, count := range currentGroup {
		if key != "nb_person" && count == nb_person {
			totalCount += 1
		}
	}
	return totalCount
}

func processAnswersEveryoneRule(inputString string, currentGroup map[string]int) map[string]int {
	for i := 0; i < len(inputString); i++ {
		char := string(inputString[i])
		_, ok := currentGroup[char]
		if ok {
			currentGroup[char] += 1
		} else {
			currentGroup[char] = 1
		}
	}
	_, ok := currentGroup["nb_person"]
	if ok {
		currentGroup["nb_person"] += 1
	} else {
		currentGroup["nb_person"] = 1
	}
	return currentGroup
}

func processAnswersAnyoneRule(inputString string, currentGroupAnswers map[string]int) map[string]int {
	for i := 0; i < len(inputString); i++ {
		char := string(inputString[i])
		_, ok := currentGroupAnswers[char]
		if !ok {
			currentGroupAnswers[char] = 1
		}
	}
	return currentGroupAnswers
}

func countGroupAnswersAnyoneRule(currentGroup map[string]int) int {
	return len(currentGroup)
}

type group struct {
	answers map[string]int
}
