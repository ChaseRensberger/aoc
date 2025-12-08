package utils

import (
	"strconv"
)

func RuneToInt(r rune) int {
	return int(r - '0')
}

func StringSliceToIntSlice(strings []string) []int {
	ints := []int{}
	for _, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ints = append(ints, n)
	}
	return ints
}
