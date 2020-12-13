package day_1

import (
	"github.com/bojeanson/advent_of_code/utils"
	"log"
)

type Pair struct {
	RightMember int
	LeftMember  int
}

type Triplet struct {
	FirstMember  int
	SecondMember int
	ThirdMember  int
}

func ResolveDay1(inputFilePath string) (Pair, Triplet) {
	lines := utils.LoadDataFromFileAndSortThem(inputFilePath)
	return FindSucceedingPair(lines), FindSucceedingTriplet(lines)
}

func FindSucceedingPair(lines []int) Pair {
	WantedValue := 2020
	var value = 0
	for len(lines) > 0 {
		value, lines, _ = utils.PopFirstElementOfSlice(lines)
		searchedValue := WantedValue - value
		pos, _ := utils.FindIntPosInSlice(searchedValue, lines)
		if pos >= 0 {
			return Pair{value, lines[pos]}
		}
	}
	log.Println("No Pair found")
	return Pair{}
}

func FindSucceedingTriplet(lines []int) Triplet {
	WantedValue := 2020
	for len(lines) > 0 {
		var firstPartOfTriplet int
		firstPartOfTriplet, lines, _ = utils.PopFirstElementOfSlice(lines)
		for len(lines) > 0 {
			var secondPartOfTriplet int
			secondPartOfTriplet, lines, _ = utils.PopFirstElementOfSlice(lines)
			searchedValue := WantedValue - firstPartOfTriplet - secondPartOfTriplet
			pos, _ := utils.FindIntPosInSlice(searchedValue, lines)
			if pos >= 0 {
				return Triplet{firstPartOfTriplet, secondPartOfTriplet, lines[pos]}
			} else {
				lines = utils.InsertInSliceAtPosition(lines, 0, secondPartOfTriplet)
				break
			}
		}
	}
	log.Println("No Triplet found")
	return Triplet{}
}
