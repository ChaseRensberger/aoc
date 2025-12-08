package y2025

import (
	// "fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Day5 struct{}

func (*Day5) SolvePartOne(input string) string {
	ans := 0
	spaceSplitInput := strings.Split(strings.TrimSpace(input), "\n\n")
	freshRanges := [][]int{}
	for _, stringRange := range strings.Split(spaceSplitInput[0], "\n") {
		rangeSides := strings.Split(stringRange, "-")
		left, err := strconv.Atoi(rangeSides[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(rangeSides[1])
		if err != nil {
			log.Fatal(err)
		}
		freshRanges = append(freshRanges, []int{left, right})
	}
	ingredientIds := strings.Split(spaceSplitInput[1], "\n")

	for _, id := range ingredientIds {
		numId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		for _, freshRange := range freshRanges {
			if numId >= freshRange[0] && numId <= freshRange[1] {
				ans++
				break
			}
		}
	}

	return strconv.Itoa(ans)
}

func (*Day5) SolvePartTwo(input string) string {
	spaceSplitInput := strings.Split(strings.TrimSpace(input), "\n\n")
	freshRanges := [][]int{}
	for _, stringRange := range strings.Split(spaceSplitInput[0], "\n") {
		rangeSides := strings.Split(stringRange, "-")
		left, err := strconv.Atoi(rangeSides[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(rangeSides[1])
		if err != nil {
			log.Fatal(err)
		}
		freshRanges = append(freshRanges, []int{left, right})
	}

	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	currentEnd := freshRanges[0][1]
	ans := freshRanges[0][1] - freshRanges[0][0] + 1
	for i := 1; i < len(freshRanges); i++ {
		start := freshRanges[i][0]
		end := freshRanges[i][1]

		if start > currentEnd+1 {
			ans += end - start + 1
			currentEnd = end
		} else if end > currentEnd {
			ans += end - currentEnd
			currentEnd = end
		}
		// fmt.Println(ans)
	}
	// fmt.Println(ans)

	return strconv.Itoa(ans)
}
