package y2025

import (
	"log"
	"strconv"
	"strings"
)

type Day1 struct{}

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
			dial = ((dial % 100) + 100) % 100
		} else {
			dial += dis
			dial %= 100
		}
		if dial == 0 {
			ans++
		}
	}
	return strconv.Itoa(ans)
}

func (*Day1) SolvePartTwo(input string) string {
	rotations := strings.Split(strings.TrimSpace(input), "\n")
	dial := 50
	ans := 0
	for _, r := range rotations {
		// fmt.Printf("dial: %v\n", dial)
		dir := r[0]
		dis, err := strconv.Atoi(r[1:])
		if err != nil {
			log.Fatal(err)
		}
		if dir == 'L' {
			dial -= dis
			if dial < 0 {
				ans += (dial + 99) / 100
			}
			dial = ((dial % 100) + 100) % 100
		} else {
			dial += dis
			if dial > 99 {
				ans += (dial / 100)
			}
			dial %= 100
		}
		if dial == 0 {
			ans++
		}
		// fmt.Printf("ans: %v\n", ans)
	}
	return strconv.Itoa(ans)
}
