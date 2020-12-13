package day_2

import (
	"github.com/bojeanson/advent_of_code/utils"
	"strconv"
	"strings"
)

func ResolveDay2(inputFilePath string) int {
	return countValidPassword(utils.LoadDataFromFile(inputFilePath))
}

func countValidPassword(lines []string) int {
	var validPasswordCount int
	for _, password := range lines {
		if isValidPassword(password) {
			validPasswordCount += 1
		}
	}
	return validPasswordCount
}

func isValidPassword(policyWithPassword string) bool {
	splitPolicyPassword := strings.Split(policyWithPassword, ":")
	policy, password := splitPolicyPassword[0], splitPolicyPassword[1]

	splitNumberOfTimeLetter := strings.Split(policy, " ")
	numberOfTime, letter := splitNumberOfTimeLetter[0], splitNumberOfTimeLetter[1]

	splitNumberOfTime := strings.Split(numberOfTime, "-")
	minNumber, _ := strconv.ParseInt(splitNumberOfTime[0], 10, 64)
	maxNumber, _ := strconv.ParseInt(splitNumberOfTime[1], 10, 64)

	letterOccurenceNumber := countLetterOccurenceNumber(letter, strings.TrimSpace(password))
	if letterOccurenceNumber >= int(minNumber) && letterOccurenceNumber <= int(maxNumber) {
		return true
	}

	return false
}

func countLetterOccurenceNumber(letter string, password string) int {
	return strings.Count(password, letter)
}
