package y2025

import "testing"

var day5Sample = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestDay5Part1Sample(t *testing.T) {
	expected := "3"
	result := (&Day5{}).SolvePartOne(day5Sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay5Part2Sample(t *testing.T) {
	expected := "53"

	result := (&Day5{}).SolvePartTwo(day5Sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
