package y2025

import (
	"log"
	"strconv"
	"strings"
)

type Day5 struct{}

func (*Day5) SolvePartOne(input string) string {
	ans := 0
	spaceSplitInput := strings.Split(strings.TrimSpace(input), "\n\n")
	freshRanges := strings.Split(spaceSplitInput[0], "\n")
	ingredientIds := strings.Split(spaceSplitInput[1], "\n")
	valid := map[int]struct{}{}

	for _, idRange := range freshRanges {
		rangeSides := strings.Split(idRange, "-")
		left, err := strconv.Atoi(rangeSides[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(rangeSides[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := left; i <= right; i++ {
			valid[i] = struct{}{}
		}
	}

	for _, id := range ingredientIds {
		numId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		if _, exists := valid[numId]; exists {
			ans++
		}
	}

	return strconv.Itoa(ans)
}

func (*Day5) SolvePartTwo(input string) string {
	return ""
}
