package y2025

import "testing"

var day3Sample = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestDay3Part1Sample(t *testing.T) {
	expected := "357"
	result := (&Day3{}).SolvePartOne(day3Sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay3Part2Sample(t *testing.T) {
	expected := "3121910778619"

	result := (&Day3{}).SolvePartTwo(day3Sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
