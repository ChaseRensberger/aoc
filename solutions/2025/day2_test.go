package y2025

import "testing"

var day2Sample = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestDay2Part1Sample(t *testing.T) {
	expected := "1227775554"
	result := (&Day2{}).SolvePartOne(day2Sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay2Part2Sample(t *testing.T) {
	expected := "4174379265"
	result := (&Day2{}).SolvePartTwo(day2Sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
