package main

import (
	"flag"
	"fmt"
	"os"

	"aoc/solutions"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env.local")
}

func main() {
	year := flag.Int("year", 0, "year")
	day := flag.Int("day", 0, "day")
	part := flag.Int("part", 0, "part (1 or 2, default: both)")
	flag.Parse()

	if *year == 0 || *day == 0 {
		fmt.Println("Usage: go run . --year 2022 --day 1 [--part 1]")
		os.Exit(1)
	}

	daySolver, err := solutions.GetDaySolver(*year, *day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input, err := loadInputFile(*year, *day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *part == 0 || *part == 1 {
		fmt.Println(daySolver.SolvePartOne(input))
	}
	if *part == 0 || *part == 2 {
		fmt.Println(daySolver.SolvePartTwo(input))
	}
}
