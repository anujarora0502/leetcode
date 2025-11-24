package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	arr := []int{2,3,5}
	target := 8

	fmt.Print(solutions.CombinationSum(arr, target))
}
