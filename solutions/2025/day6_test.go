package y2025

import "testing"

var day6Sample = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func TestDay6Part1Sample(t *testing.T) {
	expected := "4277556"
	result := (&Day6{}).SolvePartOne(day6Sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay6Part2Sample(t *testing.T) {
	expected := "6121910778619"

	result := (&Day6{}).SolvePartTwo(day6Sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
