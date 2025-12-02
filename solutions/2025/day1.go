package y2025

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Day1 struct{}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func (*Day1) SolvePartOne(input string) string {
	rotations := strings.Split(strings.TrimSpace(input), "\n")
	dial := 50
	ans := 0
	for _, r := range rotations {
		dir := r[0]
		dis, err := strconv.Atoi(r[1:])
		if err != nil {
			log.Fatal(err)
		}
		if dir == 'L' {
			dial -= dis
		} else {
			dial += dis
		}
		dial = mod(dial, 100)
		if dial == 0 {
			ans++
		}
	}
	return strconv.Itoa(ans)
}

func (*Day1) SolvePartTwo(input string) string {
	var newPos int
	var lasPos int
	rotations := strings.Split(strings.TrimSpace(input), "\n")
	dial := 50
	ans := 0
	for _, r := range rotations {
		newPos = dial
		dir := r[0]
		dis, err := strconv.Atoi(r[1:])
		if err != nil {
			log.Fatal(err)
		}
		if dir == 'L' {
			newPos -= dis
		} else {
			newPos += dis
		}
		lasPos = dial
		dial = mod(newPos, 100)
		if dial == 0 {
			ans++
		}
		if newPos > 99 {
			ans += (newPos / 100)
			if dial == 0 {
				ans--
			}
		} else if newPos < 0 {
			ans += (abs(newPos) + 99) / 100
			if lasPos == 0 {
				ans--
			}
		}
	}
	return strconv.Itoa(ans)
}
