package y2025

import "testing"

var sample = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestDay1Part1Sample(t *testing.T) {
	expected := "3"
	result := (&Day1{}).SolvePartOne(sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay1Part2Sample(t *testing.T) {
	expected := "6"
	result := (&Day1{}).SolvePartTwo(sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
