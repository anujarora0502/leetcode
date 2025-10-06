package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	var first solutions.ListNode
	first.Val = 1
	var second solutions.ListNode
	second.Val = 1 
	var third solutions.ListNode
	third.Val = 88
	first.Next = &second
	second.Next = &third
	fmt.Println(solutions.DeleteDuplicates(&first))
}
