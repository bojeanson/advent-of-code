package day_7

import (
	"github.com/bojeanson/advent_of_code/utils"
	"regexp"
	"strconv"
	"strings"
)

func ResolveDay7(part string, filename string) int {
	var bags []bag
	for inputString := range utils.StringIteratorFromFile(filename) {
		bags = processBags(inputString, bags)
	}
	var nbBagsAllowingShinyBag int
	if part == "1" {
		nbBagsAllowingShinyBag = len(getBagsAllowingSpecificBag(bags, "shiny gold", []bag{}))
	} else {
		nbBagsAllowingShinyBag = getBagsRequiredInsideSpecificBag(bags, "shiny gold", 0)
	}
	return nbBagsAllowingShinyBag
}

func getBagsRequiredInsideSpecificBag(bags []bag, searchedBagName string, nbBagsAllowingShinyBag int) int {
	for _, currentBag := range bags {
		if currentBag.bagName == searchedBagName {
			if currentBag.authorizedBags == nil {
				return nbBagsAllowingShinyBag
			}
			for _, authorizedBag := range currentBag.authorizedBags {
				nbBagsAllowingShinyBag += authorizedBag.nbBag * getBagsRequiredInsideSpecificBag(bags, authorizedBag.bagName, currentBag.nbBag)
			}
		}
	}
	return nbBagsAllowingShinyBag
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

func processBags(inputString string, bags []bag) []bag {
	indexContains := strings.Index(inputString, "bags")
	containerBag := bag{
		bagName: strings.TrimSpace(inputString[:indexContains]),
		nbBag:   1,
	}
	containedBags := inputString[indexContains+len("contain"):]
	containerBag.authorizedBags = parseAuthorizedBags(containedBags)
	bags = append(bags, containerBag)
	return bags
}

func parseAuthorizedBags(inputString string) []bag {
	re := regexp.MustCompile(`(\d{1}) (\w{3,10} \w{3,10}) bags?[,.]|no other bags\.`)
	matched := re.FindAllSubmatch([]byte(inputString), -1)
	var authorizedBags []bag
	for _, match := range matched {
		wholeMatch := string(match[0])
		if wholeMatch != "no other bags." {
			intAsInt64, _ := strconv.ParseInt(string(match[1]), 10, 64)
			authorizedBags = append(authorizedBags, bag{bagName: string(match[2]), nbBag: int(intAsInt64)})
		}
	}
	return authorizedBags
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
	nbBag          int
	authorizedBags []bag
}
