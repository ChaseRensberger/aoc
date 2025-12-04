package y2025

import (
	// "fmt"
	"strconv"
	"strings"
)

type Day4 struct{}

var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func count(room [][]rune, i int, j int) int {
	count := 0
	for _, dir := range dirs {
		yToCheck := i + dir[0]
		xToCheck := j + dir[1]
		if yToCheck < 0 || yToCheck > len(room)-1 || xToCheck < 0 || xToCheck > len(room[0])-1 {
			continue
		}
		if room[yToCheck][xToCheck] == '@' {
			count++
		}
	}
	return count
}

func hasLessThanFourAdj(room [][]rune, i int, j int) bool {
	target := 4
	count := 0
	for _, dir := range dirs {
		yToCheck := i + dir[0]
		xToCheck := j + dir[1]
		if yToCheck < 0 || yToCheck > len(room)-1 || xToCheck < 0 || xToCheck > len(room[0])-1 {
			continue
		}
		if room[yToCheck][xToCheck] == '@' {
			count++
			if count >= target {
				return false
			}
		}
	}
	return true
}

func (*Day4) SolvePartOne(input string) string {
	ans := 0
	rows := strings.Split(strings.TrimSpace(input), "\n")
	var room [][]rune
	for i, row := range rows {
		room = append(room, []rune{})
		for _, col := range row {
			room[i] = append(room[i], col)
		}
	}

	for i, row := range room {
		for j := range row {
			// fmt.Printf("count for %v,%v: %v\n", i, j, count(room, i, j))
			// fmt.Printf("%c", room[i][j])
			if room[i][j] == '@' && hasLessThanFourAdj(room, i, j) {
				ans++
			}
		}
	}
	return strconv.Itoa(ans)
}

func (*Day4) SolvePartTwo(input string) string {
	ans := 0
	rows := strings.Split(strings.TrimSpace(input), "\n")
	room := [][]rune{}
	for i, row := range rows {
		room = append(room, []rune{})
		for _, col := range row {
			room[i] = append(room[i], col)
		}
	}

	canIterate := true
	for canIterate {
		nextRoom := [][]rune{}
		canIterate = false
		for i, row := range room {
			nextRoom = append(nextRoom, []rune{})
			for j := range row {
				// fmt.Printf("count for %v,%v: %v\n", i, j, count(room, i, j))
				// fmt.Printf("%c", room[i][j])
				if room[i][j] == '@' && hasLessThanFourAdj(room, i, j) {
					nextRoom[i] = append(nextRoom[i], '.')
					canIterate = true
					ans++
				} else {
					nextRoom[i] = append(nextRoom[i], room[i][j])
				}
			}
		}
		room = nextRoom
	}
	return strconv.Itoa(ans)
}
