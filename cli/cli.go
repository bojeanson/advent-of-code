package cli

import (
	"fmt"
	"github.com/bojeanson/advent_of_code/day_1"
	"github.com/bojeanson/advent_of_code/day_10"
	"github.com/bojeanson/advent_of_code/day_11"
	"github.com/bojeanson/advent_of_code/day_12"
	"github.com/bojeanson/advent_of_code/day_13"
	"github.com/bojeanson/advent_of_code/day_14"
	"github.com/bojeanson/advent_of_code/day_15"
	"github.com/bojeanson/advent_of_code/day_16"
	"github.com/bojeanson/advent_of_code/day_17"
	"github.com/bojeanson/advent_of_code/day_18"
	"github.com/bojeanson/advent_of_code/day_19"
	"github.com/bojeanson/advent_of_code/day_2"
	"github.com/bojeanson/advent_of_code/day_20"
	"github.com/bojeanson/advent_of_code/day_3"
	"github.com/bojeanson/advent_of_code/day_4"
	"github.com/bojeanson/advent_of_code/day_5"
	"github.com/bojeanson/advent_of_code/day_6"
	"github.com/bojeanson/advent_of_code/day_7"
	"github.com/bojeanson/advent_of_code/day_8"
	"github.com/bojeanson/advent_of_code/day_9"
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

	day7           = app.Command("day_7", "resolve day_7 problem")
	day7Action     = day7.Action(uday7)
	part1Or2Day7   = day7.Arg("part", "which part of day_7 problem to resolve").Default("1").String()
	inputFilePath7 = day7.Arg("input_file", "the path to the file needed for day_7 problem").Default("./day_7/input.txt").String()

	day8           = app.Command("day_8", "resolve day_8 problem")
	day8Action     = day8.Action(uday8)
	part1Or2Day8   = day8.Arg("part", "which part of day_8 problem to resolve").Default("1").String()
	inputFilePath8 = day8.Arg("input_file", "the path to the file needed for day_8 problem").Default("./day_8/input.txt").String()

	day9           = app.Command("day_9", "resolve day_9 problem")
	day9Action     = day9.Action(uday9)
	part1Or2Day9   = day9.Arg("part", "which part of day_9 problem to resolve").Default("1").String()
	inputFilePath9 = day9.Arg("input_file", "the path to the file needed for day_9 problem").Default("./day_9/input.txt").String()

	day10           = app.Command("day_10", "resolve day_10 problem")
	day10Action     = day10.Action(uday10)
	part1Or2Day10   = day10.Arg("part", "which part of day_10 problem to resolve").Default("1").String()
	inputFilePath10 = day10.Arg("input_file", "the path to the file needed for day_10 problem").Default("./day_10/input.txt").String()

	day11           = app.Command("day_11", "resolve day_11 problem")
	day11Action     = day11.Action(uday11)
	part1Or2Day11   = day11.Arg("part", "which part of day_11 problem to resolve").Default("1").String()
	inputFilePath11 = day11.Arg("input_file", "the path to the file needed for day_11 problem").Default("./day_11/input.txt").String()

	day12           = app.Command("day_12", "resolve day_12 problem")
	day12Action     = day12.Action(uday12)
	part1Or2Day12   = day12.Arg("part", "which part of day_12 problem to resolve").Default("1").String()
	inputFilePath12 = day12.Arg("input_file", "the path to the file needed for day_12 problem").Default("./day_12/input.txt").String()

	day13           = app.Command("day_13", "resolve day_13 problem")
	day13Action     = day13.Action(uday13)
	part1Or2Day13   = day13.Arg("part", "which part of day_13 problem to resolve").Default("1").String()
	inputFilePath13 = day13.Arg("input_file", "the path to the file needed for day_13 problem").Default("./day_13/input.txt").String()

	day14           = app.Command("day_14", "resolve day_14 problem")
	day14Action     = day14.Action(uday14)
	part1Or2Day14   = day14.Arg("part", "which part of day_14 problem to resolve").Default("1").String()
	inputFilePath14 = day14.Arg("input_file", "the path to the file needed for day_14 problem").Default("./day_14/input.txt").String()

	day15           = app.Command("day_15", "resolve day_15 problem")
	day15Action     = day15.Action(uday15)
	part1Or2Day15   = day15.Arg("part", "which part of day_15 problem to resolve").Default("1").String()
	inputFilePath15 = day15.Arg("input_file", "the path to the file needed for day_15 problem").Default("./day_15/input.txt").String()

	day16           = app.Command("day_16", "resolve day_16 problem")
	day16Action     = day16.Action(uday16)
	part1Or2Day16   = day16.Arg("part", "which part of day_16 problem to resolve").Default("1").String()
	inputFilePath16 = day16.Arg("input_file", "the path to the file needed for day_16 problem").Default("./day_16/input.txt").String()

	day17           = app.Command("day_17", "resolve day_17 problem")
	day17Action     = day17.Action(uday17)
	part1Or2Day17   = day17.Arg("part", "which part of day_17 problem to resolve").Default("1").String()
	inputFilePath17 = day17.Arg("input_file", "the path to the file needed for day_17 problem").Default("./day_17/input.txt").String()

	day18           = app.Command("day_18", "resolve day_18 problem")
	day18Action     = day18.Action(uday18)
	part1Or2Day18   = day18.Arg("part", "which part of day_18 problem to resolve").Default("1").String()
	inputFilePath18 = day18.Arg("input_file", "the path to the file needed for day_18 problem").Default("./day_18/input.txt").String()

	day19           = app.Command("day_19", "resolve day_19 problem")
	day19Action     = day19.Action(uday19)
	part1Or2Day19   = day19.Arg("part", "which part of day_19 problem to resolve").Default("1").String()
	inputFilePath19 = day19.Arg("input_file", "the path to the file needed for day_19 problem").Default("./day_19/input.txt").String()

	day20           = app.Command("day_20", "resolve day_20 problem")
	day20Action     = day20.Action(uday20)
	part1Or2Day20   = day20.Arg("part", "which part of day_20 problem to resolve").Default("1").String()
	inputFilePath20 = day20.Arg("input_file", "the path to the file needed for day_20 problem").Default("./day_20/input.txt").String()
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
	fmt.Println("mon id de siège :", validPassports)
	return nil
}

func uday6(c *kingpin.ParseContext) error {
	nbAnswers := day_6.ResolveDay6(*part1Or2Day6, *inputFilePath6)
	fmt.Println("nombre de questions répondues :", nbAnswers)
	return nil
}

func uday7(c *kingpin.ParseContext) error {
	bagAllowingShinyGolgBagNumber := day_7.ResolveDay7(*part1Or2Day7, *inputFilePath7)
	fmt.Println("nombre de sacs pouvant contenir mon sac :", bagAllowingShinyGolgBagNumber)
	return nil
}

func uday8(c *kingpin.ParseContext) error {
	acc := day_8.ResolveDay8(*part1Or2Day8, *inputFilePath8)
	fmt.Println("Valeur de l'accumulateur :", acc)
	return nil
}

func uday9(c *kingpin.ParseContext) error {
	invalidNumber := day_9.ResolveDay9(*part1Or2Day9, *inputFilePath9)
	fmt.Println("nombre invalide :", invalidNumber)
	return nil
}

func uday10(c *kingpin.ParseContext) error {
	_ = day_10.ResolveDay10(*part1Or2Day10, *inputFilePath10)
	fmt.Println("")
	return nil
}

func uday11(c *kingpin.ParseContext) error {
	_ = day_11.ResolveDay11(*part1Or2Day11, *inputFilePath11)
	fmt.Println("")
	return nil
}

func uday12(c *kingpin.ParseContext) error {
	_ = day_12.ResolveDay12(*part1Or2Day12, *inputFilePath12)
	fmt.Println("")
	return nil
}

func uday13(c *kingpin.ParseContext) error {
	_ = day_13.ResolveDay13(*part1Or2Day13, *inputFilePath13)
	fmt.Println("")
	return nil
}

func uday14(c *kingpin.ParseContext) error {
	_ = day_14.ResolveDay14(*part1Or2Day14, *inputFilePath14)
	fmt.Println("")
	return nil
}

func uday15(c *kingpin.ParseContext) error {
	_ = day_15.ResolveDay15(*part1Or2Day15, *inputFilePath15)
	fmt.Println("")
	return nil
}

func uday16(c *kingpin.ParseContext) error {
	_ = day_16.ResolveDay16(*part1Or2Day16, *inputFilePath16)
	fmt.Println("")
	return nil
}

func uday17(c *kingpin.ParseContext) error {
	_ = day_17.ResolveDay17(*part1Or2Day17, *inputFilePath17)
	fmt.Println("")
	return nil
}

func uday18(c *kingpin.ParseContext) error {
	_ = day_18.ResolveDay18(*part1Or2Day18, *inputFilePath18)
	fmt.Println("")
	return nil
}

func uday19(c *kingpin.ParseContext) error {
	_ = day_19.ResolveDay19(*part1Or2Day19, *inputFilePath19)
	fmt.Println("")
	return nil
}

func uday20(c *kingpin.ParseContext) error {
	_ = day_20.ResolveDay20(*part1Or2Day20, *inputFilePath20)
	fmt.Println("")
	return nil
}
