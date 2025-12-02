package y2025

import (
	"log"
	"strconv"
	"strings"
)

type Day2 struct{}

func isValid(num int) bool {
	numString := strconv.Itoa(num)
	if len(numString)%2 != 0 {
		return true
	}
	o := 0
	c := len(numString) / 2
	for c < len(numString) {
		if numString[o] != numString[c] {
			return true
		}
		o++
		c++
	}

	return false
}

func (*Day2) SolvePartOne(input string) string {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	ans := 0
	for _, r := range ranges {
		unitRange := strings.Split(r, "-")
		l, err := strconv.Atoi(unitRange[0])
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(unitRange[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := l; i <= r; i++ {
			if !isValid(i) {
				ans += i
			}
		}
	}
	return strconv.Itoa(ans)
}

func isValidPart2(num int) bool {
	numString := strconv.Itoa(num)
	if len(numString) < 2 {
		return true
	}
	doubledString := numString + numString
	return !strings.Contains(doubledString[1:len(doubledString)-1], numString)
}

func (*Day2) SolvePartTwo(input string) string {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	ans := 0
	for _, r := range ranges {
		unitRange := strings.Split(r, "-")
		l, err := strconv.Atoi(unitRange[0])
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(unitRange[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := l; i <= r; i++ {
			if !isValidPart2(i) {
				ans += i
			}
		}
	}
	return strconv.Itoa(ans)
}
