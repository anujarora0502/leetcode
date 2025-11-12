package solutions

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SumNumbers(root *TreeNode) int {
	var result int
	sumNumHelper(root, "", &result)
	return result
}

func sumNumHelper(root *TreeNode, numSoFar string, result *int) {
	fmt.Println("num-so-far: ", numSoFar)
	fmt.Println("result: ", result)
	if root.Left == nil && root.Right == nil {
		r, _ := strconv.Atoi(numSoFar + strconv.Itoa(root.Val))
		*result += r
		return
	}

	if root.Right == nil {
		sumNumHelper(root.Left, numSoFar+strconv.Itoa(root.Val), result)
		return
	}

	if root.Left == nil {
		sumNumHelper(root.Right, numSoFar+strconv.Itoa(root.Val), result)
		return
	}

	sumNumHelper(root.Left, numSoFar+strconv.Itoa(root.Val), result)
	sumNumHelper(root.Right, numSoFar+strconv.Itoa(root.Val), result)
}
