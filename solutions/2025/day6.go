package y2025

import (
	"aoc/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Day6 struct{}

func (*Day6) SolvePartOne(input string) string {
	ans := 0
	newLineSplitInput := strings.Split(strings.TrimSpace(input), "\n")
	table := [][]int{}
	for i := 0; i < len(newLineSplitInput)-1; i++ {
		spaceSplitRow := strings.Split(strings.TrimSpace(newLineSplitInput[i]), " ")
		table = append(table, utils.StringSliceToIntSlice(spaceSplitRow))
	}

	ops := strings.Split(newLineSplitInput[len(newLineSplitInput)-1], " ")
	ops = slices.DeleteFunc(ops, func(s string) bool {
		return s == ""
	})
	m := len(table[0])
	n := len(table)
	var rr int
	for j := range m {
		switch ops[j] {
		case "+":
			rr = 0
		case "*":
			rr = 1
		}
		for i := range n {
			switch ops[j] {
			case "+":
				rr += table[i][j]
			case "*":
				rr *= table[i][j]
			}
		}
		ans += rr
	}

	return strconv.Itoa(ans)
}

func computeColumn(s string, c int, n int, op string) int {
	var res int
	switch op {
	case "+":
		res = 0
	case "*":
		res = 1
	}
	for i := c; i < len(s); i += n {
		val, err := strconv.Atoi(string(s[i]))
		if err != nil {
			log.Fatal(err)
		}
		switch op {
		case "+":
			res += val
		case "*":
			res *= val
		}
	}
	return res
}

func (*Day6) SolvePartTwo(input string) string {
	ans := 0
	newLineSplitInput := strings.Split(input, "\n")
	n := len(newLineSplitInput[0])
	ops := slices.DeleteFunc(strings.Split(newLineSplitInput[len(newLineSplitInput)-1], " "), func(s string) bool {
		return s == ""
	})

	for i := range ops {

	}

	return strconv.Itoa(ans)
}
