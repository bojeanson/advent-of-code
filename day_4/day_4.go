package day_4

import (
	"github.com/bojeanson/advent_of_code/utils"
	"regexp"
	"strconv"
	"strings"
)

func ResolveDay4(filename string) int {
	validPassports := 0
	currentPassport := passport{}
	for inputString := range utils.StringIteratorFromFile(filename) {
		switch inputString {
		case "":
			if inputString == "" {
				if checkPassportValidity(currentPassport) {
					validPassports += 1
				}
				currentPassport = passport{}
			}
		default:
			currentPassport = checkFields(inputString, currentPassport)
		}
	}
	if checkPassportValidity(currentPassport) {
		validPassports += 1
	}
	return validPassports
}

func checkPassportValidity(currentPassport passport) bool {
	if currentPassport.byr && currentPassport.ecl && currentPassport.eyr && currentPassport.hcl && currentPassport.hgt && currentPassport.iyr && currentPassport.pid {
		return true
	}
	return false
}

func checkFields(inputString string, currentPassport passport) passport {
	if strings.Contains(inputString, "byr") {
		currentPassport = validateByr(inputString, currentPassport)
	}
	if strings.Contains(inputString, "cid") {
		currentPassport.cid = true
	}
	if strings.Contains(inputString, "ecl") {
		currentPassport = validateEcl(inputString, currentPassport)
	}
	if strings.Contains(inputString, "eyr") {
		currentPassport = validateEyr(inputString, currentPassport)
	}
	if strings.Contains(inputString, "hcl") {
		currentPassport = validateHcl(inputString, currentPassport)
	}
	if strings.Contains(inputString, "hgt") {
		currentPassport = validateHgt(inputString, currentPassport)

	}
	if strings.Contains(inputString, "iyr") {
		currentPassport = validateIyr(inputString, currentPassport)

	}
	if strings.Contains(inputString, "pid") {
		currentPassport.pid = validPid(inputString)
	}
	return currentPassport
}

func validateIyr(inputString string, currentPassport passport) passport {
	minBound, maxBound := 2010, 2020
	indexIyr := strings.Index(inputString, "iyr")
	iyr := extractDateAtIndex(inputString, indexIyr)
	currentPassport.iyr = validDate(iyr, minBound, maxBound, currentPassport)
	return currentPassport
}

func validateEyr(inputString string, currentPassport passport) passport {
	minBound, maxBound := 2020, 2030
	indexByr := strings.Index(inputString, "eyr")
	eyr := extractDateAtIndex(inputString, indexByr)
	currentPassport.eyr = validDate(eyr, minBound, maxBound, currentPassport)
	return currentPassport
}

func validateByr(inputString string, currentPassport passport) passport {
	minBound, maxBound := 1920, 2002
	indexByr := strings.Index(inputString, "byr")
	byr := extractDateAtIndex(inputString, indexByr)
	currentPassport.byr = validDate(byr, minBound, maxBound, currentPassport)
	return currentPassport
}

func extractDateAtIndex(inputString string, index int) int {
	date, _ := strconv.ParseInt(inputString[index+4:index+4+4], 10, 64)
	return int(date)
}

func validDate(date int, minBound int, maxBound int, currentPassport passport) bool {
	if date >= minBound && date <= maxBound {
		return true
	}
	return false
}

func validateHgt(inputString string, currentPassport passport) passport {
	matched, _ := regexp.Match("hgt:((15[0-9]|1[6-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)", []byte(inputString))
	currentPassport.hgt = matched
	return currentPassport
}

func validateHcl(inputString string, currentPassport passport) passport {
	matched, _ := regexp.Match("hcl:#(\\d|[a-f]){6}?", []byte(inputString))
	currentPassport.hcl = matched
	return currentPassport
}

func validateEcl(inputString string, currentPassport passport) passport {
	matched, _ := regexp.Match("ecl:(amb|blu|brn|gry|grn|hzl|oth)", []byte(inputString))
	currentPassport.ecl = matched
	return currentPassport
}

func validPid(inputString string) bool {
	matched, _ := regexp.Match("pid:\\d{9}($|(^\\d|^\\W|\\s))", []byte(inputString))
	return matched
}

type passport struct {
	byr bool
	cid bool
	ecl bool
	eyr bool
	hcl bool
	hgt bool
	iyr bool
	pid bool
}
