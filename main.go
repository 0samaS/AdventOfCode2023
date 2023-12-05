package main

import(
	"github.com/0samaS/AdventOfCode2023/sols"
	"github.com/fatih/color"
	"os"
)

func main() {
	part := os.Args[1]
	yellow := color.New(color.FgYellow)
	yellow.Printf("Solution for part %s is: %d\n", part, sols.Day5(part))
}
