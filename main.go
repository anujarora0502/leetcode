package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	heightMap := [][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}

	fmt.Println(solutions.TrapRainWater(heightMap))
}
