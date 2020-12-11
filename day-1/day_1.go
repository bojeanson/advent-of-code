package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := loadDataFromFileAndSortThem("input.txt")
	pair1, pair2 := findSucceedingPair(lines)
	triplet1, triplet2, triplet3 := findSucceedingTriplet(lines)
	fmt.Println("(pair1,pair2)=(", pair1, ",", pair2, ")")
	fmt.Println("pair1+pair2=", pair1+pair2)
	fmt.Println("pair1xpair2=", pair1*pair2)
	fmt.Println("(triplet1,triplet2,triplet3)=(", triplet1, ",", triplet2, ",", triplet3, ")")
	fmt.Println("triplet1+triplet2+triplet3=", triplet1+triplet2+triplet3)
	fmt.Println("triplet1xtriplet2=", triplet1*triplet2*triplet3)
}

func loadDataFromFileAndSortThem(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputIntAsInt64, err := strconv.ParseInt(scanner.Text(), 10, 64)
		inputInt := int(inputIntAsInt64)
		if err != nil {
			log.Fatal(err)
		}

		switch {
		case lines == nil || inputInt > lines[len(lines)-1]:
			lines = append(lines, inputInt)
		default:
			for i, _ := range lines {
				if inputInt <= lines[i] {
					lines = insertAtPosition(lines, i, inputInt)
					break
				}
			}
		}
	}
	return lines
}

func insertAtPosition(lines []int, i int, inputInt int) []int {
	lines = append(lines, 0)
	copy(lines[i+1:], lines[i:])
	lines[i] = inputInt
	return lines
}

func findSucceedingPair(lines []int) (int, int) {
	WantedValue := 2020
	var value = 0
	for len(lines) > 0 {
		value, lines = popFirstElementOfSlice(lines)
		searchedValue := WantedValue - value
		pos := findIntPosInSlice(searchedValue, lines)
		if pos > 0 {
			return value, lines[pos]
		}
	}
	log.Fatal("No pair found")
	return 0, 0
}

func popFirstElementOfSlice(lines []int) (int, []int) {
	var value int
	value, lines = lines[0], lines[1:]
	return value, lines
}

func findSucceedingTriplet(lines []int) (int, int, int) {
	WantedValue := 2020
	for len(lines) > 0 {
		var firstPartOfTriplet int
		firstPartOfTriplet, lines = popFirstElementOfSlice(lines)
		for len(lines) > 0 {
			var secondPartOfTriplet int
			secondPartOfTriplet, lines = popFirstElementOfSlice(lines)
			searchedValue := WantedValue - firstPartOfTriplet - secondPartOfTriplet
			pos := findIntPosInSlice(searchedValue, lines)
			if pos > 0 {
				return firstPartOfTriplet, secondPartOfTriplet, lines[pos]
			} else {
				lines = insertAtPosition(lines, 0, secondPartOfTriplet)
				break
			}
		}
	}
	log.Fatal("No triplet found")
	return 0, 0, 0
}


func findIntPosInSlice(searchedValue int, list []int) int {
	for i, element := range list {
		if element == searchedValue {
			return i
		}
	}
	return 0
}