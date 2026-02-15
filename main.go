package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	list := []int{10, 20, 30, 40, 35, 45, 42, 55, 30, 15}

	pq := solutions.Initialize()

	for _, e := range list {
		pq.Add(e)
	}

	sortedArray := make([]int, 0)

	for pq.Size() != 0 {
		pq.PrintQueue()
		sortedArray = append(sortedArray, pq.Remove())
	}

	fmt.Println("Sorted Array: ", sortedArray)
}
