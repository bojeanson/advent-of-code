package cli

import (
	"fmt"
	"github.com/bojeanson/advent_of_code/day_1"
	"github.com/bojeanson/advent_of_code/day_2"
	"github.com/bojeanson/advent_of_code/day_3"
	"github.com/bojeanson/advent_of_code/day_4"
	"github.com/bojeanson/advent_of_code/day_5"
	"github.com/bojeanson/advent_of_code/day_6"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("advent_of_code", "Resolve advent-of-code problem for each day")

	day1           = app.Command("day_1", "resolve day_1 problem")
	day1Action     = day1.Action(uday1)
	inputFilePath1 = day1.Arg("input_file", "the path to the file needed for day_1 problem").Default("./day_1/input.txt").String()

	day2           = app.Command("day_2", "resolve day_2 problem")
	day2Action     = day2.Action(uday2)
	part1Or2Day2   = day2.Arg("part", "which part of day_2 problem to resolve").Default("1").String()
	inputFilePath2 = day2.Arg("input_file", "the path to the file needed for day_2 problem").Default("./day_2/input.txt").String()

	day3           = app.Command("day_3", "resolve day_3 problem")
	day3Action     = day3.Action(uday3)
	inputFilePath3 = day3.Arg("input_file", "the path to the file needed for day_3 problem").Default("./day_3/input.txt").String()

	day4           = app.Command("day_4", "resolve day_4 problem")
	day4Action     = day4.Action(uday4)
	inputFilePath4 = day4.Arg("input_file", "the path to the file needed for day_4 problem").Default("./day_4/input.txt").String()

	day5           = app.Command("day_5", "resolve day_5 problem")
	day5Action     = day5.Action(uday5)
	inputFilePath5 = day5.Arg("input_file", "the path to the file needed for day_5 problem").Default("./day_5/input.txt").String()

	day6           = app.Command("day_6", "resolve day_6 problem")
	day6Action     = day6.Action(uday6)
	part1Or2Day6   = day6.Arg("part", "which part of day_6 problem to resolve").Default("1").String()
	inputFilePath6 = day6.Arg("input_file", "the path to the file needed for day_6 problem").Default("./day_6/input.txt").String()
)

func Parse(args []string) {
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func uday1(c *kingpin.ParseContext) error {
	succeedingPair, succeedingTriplet := day_1.ResolveDay1(*inputFilePath1)
	fmt.Println("pair : (", succeedingPair.RightMember, ",", succeedingPair.LeftMember, ")")
	fmt.Println("somme de la pair :", succeedingPair.RightMember+succeedingPair.LeftMember)
	fmt.Println("produit de la pair :", succeedingPair.RightMember*succeedingPair.LeftMember)
	fmt.Println("triplet : (", succeedingTriplet.FirstMember, ",", succeedingTriplet.SecondMember, ",", succeedingTriplet.ThirdMember, ")")
	fmt.Println("somme du triplet :", succeedingTriplet.FirstMember+succeedingTriplet.SecondMember+succeedingTriplet.ThirdMember)
	fmt.Println("produit du triplet :", succeedingTriplet.FirstMember*succeedingTriplet.SecondMember*succeedingTriplet.ThirdMember)
	return nil
}

func uday2(c *kingpin.ParseContext) error {
	validPasswordCount := day_2.ResolveDay2(*part1Or2Day2, *inputFilePath2)
	fmt.Println("nombre de password valides :", validPasswordCount)
	return nil
}

func uday3(c *kingpin.ParseContext) error {
	countTree := day_3.ResolveDay3(*inputFilePath3)
	fmt.Println("nombre d'arbres :", countTree)
	return nil
}

func uday4(c *kingpin.ParseContext) error {
	validPassports := day_4.ResolveDay4(*inputFilePath4)
	fmt.Println("nombre de passeports valides :", validPassports)
	return nil
}

func uday5(c *kingpin.ParseContext) error {
	validPassports := day_5.ResolveDay5(*inputFilePath5)
	fmt.Println("My seat ID :", validPassports)
	return nil
}

func uday6(c *kingpin.ParseContext) error {
	nbAnswers := day_6.ResolveDay6(*part1Or2Day6, *inputFilePath6)
	fmt.Println("Nombre de réponses répondues :", nbAnswers)
	return nil
}
