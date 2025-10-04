package solutions

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if heightLeft(root) == heightRight(root) {
		return int(math.Pow(float64(2), float64(heightLeft(root)))) - 1
	}

	return CountNodes(root.Left) + CountNodes(root.Right) + 1
}

func heightLeft(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return heightLeft(root.Left) + 1
}

func heightRight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return heightRight(root.Right) + 1
}

func CreateCompleteBinaryTree() *TreeNode {
	// Create nodes for level 5
	nodes := make([]TreeNode, 31)
	for i := range nodes {
		nodes[i].Val = i + 1
	}
	
	// Build complete binary tree structure
	// Level 1
	root := &nodes[0]
	
	// Level 2
	root.Left = &nodes[1]    // 2
	root.Right = &nodes[2]   // 3
	
	// Level 3
	root.Left.Left = &nodes[3]     // 4
	root.Left.Right = &nodes[4]    // 5
	root.Right.Left = &nodes[5]    // 6  
	root.Right.Right = &nodes[6]   // 7
	
	// Level 4
	root.Left.Left.Left = &nodes[7]     // 8
	root.Left.Left.Right = &nodes[8]    // 9
	root.Left.Right.Left = &nodes[9]    // 10
	root.Left.Right.Right = &nodes[10]  // 11
	root.Right.Left.Left = &nodes[11]   // 12
	root.Right.Left.Right = &nodes[12]  // 13
	root.Right.Right.Left = &nodes[13]  // 14
	root.Right.Right.Right = &nodes[14] // 15
	
	// Level 5
	root.Left.Left.Left.Left = &nodes[15]     // 16
	root.Left.Left.Left.Right = &nodes[16]    // 17
	root.Left.Left.Right.Left = &nodes[17]    // 18
	root.Left.Left.Right.Right = &nodes[18]   // 19
	root.Left.Right.Left.Left = &nodes[19]    // 20
	root.Left.Right.Left.Right = &nodes[20]   // 21
	root.Left.Right.Right.Left = &nodes[21]   // 22
	root.Left.Right.Right.Right = &nodes[22]  // 23
	root.Right.Left.Left.Left = &nodes[23]    // 24
	root.Right.Left.Left.Right = &nodes[24]   // 25
	root.Right.Left.Right.Left = &nodes[25]   // 26
	root.Right.Left.Right.Right = &nodes[26]  // 27
	root.Right.Right.Left.Left = &nodes[27]   // 28
	root.Right.Right.Left.Right = &nodes[28]  // 29
	root.Right.Right.Right.Left = &nodes[29]  // 30
	root.Right.Right.Right.Right = &nodes[30] // 31
	
	return root
}

