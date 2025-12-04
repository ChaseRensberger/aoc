package y2025

import "testing"

var day4Sample = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestDay4Part1Sample(t *testing.T) {
	expected := "13"
	result := (&Day4{}).SolvePartOne(day4Sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay4Part2Sample(t *testing.T) {
	expected := "43"

	result := (&Day4{}).SolvePartTwo(day4Sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
