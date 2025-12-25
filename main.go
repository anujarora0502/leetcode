package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	intervals := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	interval := []int{4, 8}

	fmt.Println(solutions.Insert(intervals, interval))
}
