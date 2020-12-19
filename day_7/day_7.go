package day_7

import (
	"github.com/bojeanson/advent_of_code/utils"
	"regexp"
	"strings"
)

func ResolveDay7(filename string) int {
	var bags []bag
	for inputString := range utils.StringIteratorFromFile(filename) {
		bags = processBags(inputString, bags)
	}
	var nbBagsAllowingShinyBag []bag
	nbBagsAllowingShinyBag = getBagsAllowingSpecificBag(bags, "shiny gold", nbBagsAllowingShinyBag)
	return len(nbBagsAllowingShinyBag)
}

func processBags(inputString string, bags []bag) []bag {
	indexContains := strings.Index(inputString, "bags")
	containerBag := bag{
		bagName: strings.TrimSpace(inputString[:indexContains]),
	}
	containedBags := inputString[indexContains+len("contain"):]
	containerBag.authorizedBags = parseAuthorizedBags(containedBags)
	bags = append(bags, containerBag)
	return bags
}

func parseAuthorizedBags(inputString string) []bag {
	re := regexp.MustCompile(`\d{1} (\w{3,10} \w{3,10}) bags?[,.]|no other bags\.`)
	matched := re.FindAllSubmatch([]byte(inputString), -1)
	var authorizedBags []bag
	for _, match := range matched {
		wholeMatch := string(match[0])
		if wholeMatch != "no other bags." {
			authorizedBags = append(authorizedBags, bag{bagName: string(match[1])})
		}
	}
	return authorizedBags
}

func getBagsAllowingSpecificBag(bags []bag, searchedBagName string, bagsAllowingSpecificBag []bag) []bag {
	for _, currentBag := range bags {
		if bagNameInBagList(searchedBagName, currentBag.authorizedBags) {
			if !bagNameInBagList(currentBag.bagName, bagsAllowingSpecificBag) {
				bagsAllowingSpecificBag = append(bagsAllowingSpecificBag, currentBag)
			}
			bagsAllowingSpecificBag = getBagsAllowingSpecificBag(bags, currentBag.bagName, bagsAllowingSpecificBag)
		}
	}
	return bagsAllowingSpecificBag
}

func bagNameInBagList(currentBagName string, bags []bag) bool {
	for _, b := range bags {
		if currentBagName == b.bagName {
			return true
		}
	}
	return false
}

type bag struct {
	bagName        string
	authorizedBags []bag
}
