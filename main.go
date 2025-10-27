package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	nums := []int{-10,-3,0,5,9}
	fmt.Println(solutions.SortedArrayToBST(nums).Left.Val)
}
