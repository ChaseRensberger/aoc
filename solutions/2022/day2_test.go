package y2022

import "testing"

func TestDay2Part1Sample(t *testing.T) {
	sample := `
	A Y
	B X
	C Z
	`
	expected := "15"
	result := (&Day2{}).SolvePartOne(sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay2Part2Sample(t *testing.T) {

	sample := `
	000
	000
	000
	`
	expected := "111"
	result := (&Day2{}).SolvePartTwo(sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
