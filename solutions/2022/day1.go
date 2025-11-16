package y2022

import (
	"container/heap"
	"strconv"
	"strings"
)

type Day1 struct{}

func (*Day1) SolvePartOne(input string) string {
	elves := strings.Split(strings.TrimSpace(input), "\n\n")
	maxCalories := 0

	for _, elfInventory := range elves {
		items := strings.Split(elfInventory, "\n")
		totalCalories := 0
		for _, item := range items {
			calories, err := strconv.Atoi(item)
			if err != nil {
				panic(err.Error())
			}
			totalCalories += calories
		}
		maxCalories = max(maxCalories, totalCalories)
	}
	return strconv.Itoa(maxCalories)
}

func (*Day1) SolvePartTwo(input string) string {
	elves := strings.Split(strings.TrimSpace(input), "\n\n")
	answer := 0

	h := &minHeap{}
	heap.Init(h)

	for _, elfInventory := range elves {
		items := strings.Split(elfInventory, "\n")
		totalCalories := 0
		for _, item := range items {
			calories, err := strconv.Atoi(item)
			if err != nil {
				panic(err.Error())
			}
			totalCalories += calories
		}
		if h.Len() < 3 {
			heap.Push(h, totalCalories)
		} else if totalCalories > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, totalCalories)
		}
	}
	for _, value := range *h {
		answer += value
	}
	return strconv.Itoa(answer)
}

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
