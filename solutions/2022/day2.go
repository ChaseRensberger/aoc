package y2022

import (
	"strconv"
	"strings"
)

var myScoreMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

type Day2 struct{}

func (*Day2) SolvePartOne(input string) string {
	lines := strings.Split(input, "\n")
	ans := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lineSlice := strings.Split(line, " ")
		ans += myScoreMap[lineSlice[1]]
		if (lineSlice[0] == "A" && lineSlice[1] == "X") || (lineSlice[0] == "B" && lineSlice[1] == "Y") || (lineSlice[0] == "C" && lineSlice[1] == "Z") {
			ans += 3
		} else if (lineSlice[0] == "A" && lineSlice[1] == "Y") || (lineSlice[0] == "B" && lineSlice[1] == "Z") || (lineSlice[0] == "C" && lineSlice[1] == "X") {
			ans += 6
		}
	}
	return strconv.Itoa(ans)
}

// X = Lose
// Y = Draw
// Z = Win
// A = Rock
// B = Paper
// C = Scissors

var myScoreMap2 = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func (*Day2) SolvePartTwo(input string) string {
	lines := strings.Split(input, "\n")
	ans := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lineSlice := strings.Split(line, " ")
		if lineSlice[1] == "Y" {
			// Need a draw
			ans += 3
			ans += myScoreMap2[lineSlice[0]]
		} else if lineSlice[1] == "Z" {
			// need a win
			ans += 6
		} else {
			// need to lose

		}

	}
	return strconv.Itoa(ans)
}
