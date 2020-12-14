package day_2

import (
	"github.com/bojeanson/advent_of_code/utils"
	"strconv"
	"strings"
)

func ResolveDay2(part string, inputFilePath string) int {
	return countValidPassword(part, utils.LoadDataFromFile(inputFilePath))
}

func countValidPassword(part string, lines []string) int {
	var policyToVerify func(password string, policyLetter string, lowerBound int, upperBound int) bool
	switch part {
	case "1":
		policyToVerify = policyOfOccurrences
	case "2":
		policyToVerify = policyOfPositions
	default:
		policyToVerify = policyOfOccurrences
	}
	var validPasswordCount int
	for _, password := range lines {
		if isValidPassword(password, policyToVerify) {
			validPasswordCount += 1
		}
	}
	return validPasswordCount
}

func isValidPassword(policyWithPassword string, policyToVerify func(password string, policyLetter string, lowerBound int, upperBound int) bool) bool {
	splitPolicyPassword := strings.Split(policyWithPassword, ":")
	policy := strings.TrimSpace(splitPolicyPassword[0])
	password := strings.TrimSpace(splitPolicyPassword[1])

	boundPatterAndPolicyLetter := strings.Split(policy, " ")
	boundPattern, policyLetter := boundPatterAndPolicyLetter[0], boundPatterAndPolicyLetter[1]

	bounds := strings.Split(boundPattern, "-")
	lowerBound, _ := strconv.ParseInt(bounds[0], 10, 64)
	upperBound, _ := strconv.ParseInt(bounds[1], 10, 64)

	return policyToVerify(password, policyLetter, int(lowerBound), int(upperBound))
}

func policyOfOccurrences(password string, policyLetter string, lowerBound int, upperBound int) bool {
	letterOccurrenceNumber := countLetterOccurenceNumber(policyLetter, password)
	if letterOccurrenceNumber >= int(lowerBound) && letterOccurrenceNumber <= int(upperBound) {
		return true
	}
	return false
}

func policyOfPositions(password string, policyLetter string, lowerBound int, upperBound int) bool {
	lowerBoundRespectsPolicy := string(password[lowerBound-1]) == policyLetter
	upperBoundRespectsPolicy := string(password[upperBound-1]) == policyLetter
	if (lowerBoundRespectsPolicy || upperBoundRespectsPolicy) && !(lowerBoundRespectsPolicy && upperBoundRespectsPolicy) {
		return true
	}
	return false
}

func countLetterOccurenceNumber(letter string, password string) int {
	return strings.Count(password, letter)
}
