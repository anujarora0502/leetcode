package solutions

import "math"

// Incorrect
func TrapRainWater(heightMap [][]int) int {

	result := 0

	rows := len(heightMap)
	cols := len(heightMap[0])

	leftMax := makeLeftMax(rows, cols, heightMap)
	rightMax := makeRightMax(rows, cols, heightMap)
	topMax := makeTopMax(rows, cols, heightMap)
	bottomMax := makeBottomMax(rows, cols, heightMap)

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			minHeight := int(math.Min(math.Min(float64(leftMax[i][j]),
				float64(rightMax[i][j])),
				math.Min(float64(topMax[i][j]),
					float64(bottomMax[i][j]))))
			if minHeight > heightMap[i][j] {
				result += minHeight - heightMap[i][j]
			}
		}
	}

	return result
}

func makeLeftMax(rows int, cols int, heightMap [][]int) [][]int {
	leftMax := makeSlice(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == 0 {
				leftMax[i][j] = heightMap[i][j]
			} else {
				if heightMap[i][j] > leftMax[i][j-1] {
					leftMax[i][j] = heightMap[i][j]
				} else {
					leftMax[i][j] = leftMax[i][j-1]
				}
			}
		}
	}

	return leftMax
}

func makeRightMax(rows int, cols int, heightMap [][]int) [][]int {
	rightMax := makeSlice(rows, cols)

	for i := 0; i < rows; i++ {
		for j := cols - 1; j >= 0; j-- {
			if j == cols-1 {
				rightMax[i][j] = heightMap[i][j]
			} else {
				if heightMap[i][j] > rightMax[i][j+1] {
					rightMax[i][j] = heightMap[i][j]
				} else {
					rightMax[i][j] = rightMax[i][j+1]
				}
			}
		}
	}

	return rightMax
}

func makeTopMax(rows int, cols int, heightMap [][]int) [][]int {
	topMax := makeSlice(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 {
				topMax[i][j] = heightMap[i][j]
			} else {
				if heightMap[i][j] > topMax[i-1][j] {
					topMax[i][j] = heightMap[i][j]
				} else {
					topMax[i][j] = topMax[i-1][j]
				}
			}
		}
	}

	return topMax
}

func makeBottomMax(rows int, cols int, heightMap [][]int) [][]int {
	bottomMax := makeSlice(rows, cols)

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i == rows-1 {
				bottomMax[i][j] = heightMap[i][j]
			} else {
				if heightMap[i][j] > bottomMax[i+1][j] {
					bottomMax[i][j] = heightMap[i][j]
				} else {
					bottomMax[i][j] = bottomMax[i+1][j]
				}
			}
		}
	}

	return bottomMax
}

func makeSlice(rows int, cols int) [][]int {
	slice := make([][]int, rows)
	for i := range slice {
		slice[i] = make([]int, cols)
	}

	return slice
}
