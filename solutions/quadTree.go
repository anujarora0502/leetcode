package solutions

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func Construct(grid [][]int) *Node {
	allOnes := isItAllOnes(grid)
	allZeros := isItAllZeros(grid)

	var val bool
	var isLeaf bool

	if allOnes {
		val = true
		isLeaf = true
	} else if allZeros {
		val = false
		isLeaf = true
	}

	if !isLeaf {
		n := len(grid)

		leftHalf := make([][]int, 0)
		for i := 0; i < n; i++ {
			leftHalf = append(leftHalf, make([]int, 0))
			for j := 0; j < n/2; j++ {
				leftHalf[i] = append(leftHalf[i], grid[i][j])
			}
		}
		rightHalf := make([][]int, 0)
		for i := 0; i < n; i++ {
			rightHalf = append(rightHalf, make([]int, 0))
			for j := n / 2; j < n; j++ {
				rightHalf[i] = append(rightHalf[i], grid[i][j])
			}
		}

		topLeftGrid := leftHalf[:n/2]
		topRightGrid := rightHalf[:n/2]
		bottomLeftGrid := leftHalf[n/2:]
		bottomRightGrid := rightHalf[n/2:]

		node := Node{val, isLeaf, Construct(topLeftGrid), Construct(topRightGrid), Construct(bottomLeftGrid), Construct(bottomRightGrid)}
		return &node
	}

	node := Node{val, isLeaf, nil, nil, nil, nil}
	return &node
}

func isItAllOnes(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func isItAllZeros(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 1 {
				return false
			}
		}
	}
	return true
}
