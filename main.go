package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	root := solutions.TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &solutions.TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right = &solutions.TreeNode{Val: 3, Left: nil, Right: nil}
  fmt.Println(solutions.SumNumbers(&root))
}
