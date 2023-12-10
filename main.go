package main

import(
	"github.com/0samaS/AdventOfCode2023/sols"
	"github.com/fatih/color"
	"os"
)

func main() {
	day := os.Args[1]
	part := os.Args[2]
	yellow := color.New(color.FgYellow)
	if day == "day1" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day1(part))
	} else if day == "day2" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day2(part))
	} else if day == "day3" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day3(part))
	} else if day == "day4" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day4(part))
	} else if day == "day5" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day5(part))
	} else if day == "day6" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day6(part))
	} else if day == "day7" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day7(part))
	} else if day == "day8" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day8(part))
	} else if day == "day9" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day9(part))
	} else if day == "day10" {
		yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day10(part))
	} else if day == "day11" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day11(part))
	} else if day == "day12" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day12(part))
	} else if day == "day13" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day13(part))
	} else if day == "day14" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day14(part))
	} else if day == "day15" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day15(part))
	} else if day == "day16" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day16(part))
	} else if day == "day17" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day17(part))
	} else if day == "day18" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day18(part))
	} else if day == "day19" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day19(part))
	} else if day == "day20" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day20(part))
	} else if day == "day21" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day21(part))
	} else if day == "day22" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day22(part))
	} else if day == "day23" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day23(part))
	} else if day == "day24" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day24(part))
	} else if day == "day25" {
		// yellow.Printf("Solution for %s %s is: %d\n", day, part, sols.Day25(part))
	}
}
