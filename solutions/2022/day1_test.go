package y2022

import "testing"

func TestDay1Part1Sample(t *testing.T) {
	sample := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	expected := "24000"
	result := (&Day1{}).SolvePartOne(sample)

	if result != expected {
		t.Errorf("Part 1: expected %s, got %s", expected, result)
	}
}

func TestDay1Part2Sample(t *testing.T) {
	sample := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	expected := "45000"
	result := (&Day1{}).SolvePartTwo(sample)

	if result != expected {
		t.Errorf("Part 2: expected %s, got %s", expected, result)
	}
}
