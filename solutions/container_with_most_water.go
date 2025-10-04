package solutions

import "math"

func MaxArea(height []int) int {
	maxArea := 0

	start := 0
	end := len(height) - 1

	for start < end {
		area := int(math.Min(float64(height[start]), float64(height[end]))) * (end - start)
		if area > maxArea {
			maxArea = area
		}

		if height[start] > height[end] {
			end--
		} else if height[start] == height[end] {
			start++
			end--
		} else {
			start++
		}
	}

	return maxArea
}
