package y2025

import (
	// "container/heap"
	// "sort"
	"strconv"
	"strings"
)

//
// type Element struct {
// 	value    int
// 	position int
// }
//
// type MinHeap []Element
//
// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
// func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Element)) }
// func (h *MinHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
//
// func KLargest(nums []int, k int) []int {
// 	h := &MinHeap{}
// 	heap.Init(h)
//
// 	for i, num := range nums {
// 		if h.Len() < k {
// 			heap.Push(h, Element{value: num, position: i})
// 		} else if num > (*h)[0].value {
// 			heap.Pop(h)
// 			heap.Push(h, Element{value: num, position: i})
// 		}
// 	}
//
// 	elements := []Element{}
// 	for h.Len() > 0 {
// 		elements = append(elements, heap.Pop(h).(Element))
// 	}
//
// 	sort.Slice(elements, func(i, j int) bool {
// 		return elements[i].position < elements[j].position
// 	})
//
// 	result := make([]int, len(elements))
// 	for i, elem := range elements {
// 		result[i] = elem.value
// 	}
//
// 	return result
// }
//
// func stringToList(s string) []int {
// 	var output []int
// 	for _, c := range s {
// 		n, err := strconv.Atoi(string(c))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		output = append(output, n)
// 	}
// 	return output
// }

func runeToInt(r rune) int {
	return int(r - '0')
}

type Day3 struct{}

func (*Day3) SolvePartOne(input string) string {
	banks := strings.Split(strings.TrimSpace(input), "\n")
	ans := 0
	for _, bank := range banks {
		first := -1
		firstPos := -1
		second := -1
		for i, battery := range bank[:len(bank)-1] {
			if runeToInt(battery) > first {
				firstPos = i
				first = runeToInt(battery)
			}
		}
		for _, battery := range bank[firstPos+1:] {
			second = max(second, runeToInt(battery))
		}
		ans += first*10 + second
	}
	return strconv.Itoa(ans)
}

func (*Day3) SolvePartTwo(input string) string {
	return ""
}
