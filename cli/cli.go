package cli

import (
	"fmt"
	"github.com/bojeanson/advent_of_code/day_1"
	"github.com/bojeanson/advent_of_code/day_2"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app            = kingpin.New("advent_of_code", "Resolve advent-of-code problem for each day")
	day1           = app.Command("day_1", "resolve day_1 problem")
	day1Action     = day1.Action(uday1)
	inputFilePath1 = day1.Arg("input_file", "the path to the file needed for day_1 problem").Default("./day_1/input.txt").String()

	day2           = app.Command("day_2", "resolve day_2 problem")
	day2Action     = day2.Action(uday2)
	inputFilePath2 = day2.Arg("input_file", "the path to the file needed for day_2 problem").Default("./day_2/input.txt").String()
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
	validPasswordCount := day_2.ResolveDay2(*inputFilePath2)
	fmt.Println("nombre de password valides :", validPasswordCount)
	return nil
}
