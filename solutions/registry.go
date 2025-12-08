package solutions

import (
	"fmt"

	y2022 "aoc/solutions/2022"
	y2025 "aoc/solutions/2025"
)

var solvers = map[int]map[int]DaySolver{
	2022: {
		1: &y2022.Day1{},
		2: &y2022.Day2{},
	},
	2025: {
		1: &y2025.Day1{},
		2: &y2025.Day2{},
		3: &y2025.Day3{},
		4: &y2025.Day4{},
		5: &y2025.Day5{},
		6: &y2025.Day6{},
	},
}

type DaySolver interface {
	SolvePartOne(string) string
	SolvePartTwo(string) string
}

func GetDaySolver(year int, day int) (DaySolver, error) {
	if yearMap, exists := solvers[year]; exists {
		if solver, exists := yearMap[day]; exists {
			return solver, nil
		}
	}
	return nil, fmt.Errorf("no solver found for year %d, day %d", year, day)
}
