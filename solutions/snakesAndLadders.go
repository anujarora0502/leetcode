package solutions

import "fmt"

func SnakesAndLadders(board [][]int) int {
	queue := make([]int, 0)

	queue = append(queue, 1)

	distance := 0
	n := len(board)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for i := 1; i <= 6; i++ {
			next := curr + i
			var row, col int
			if next%n == 0 {
				row = n - (next / n)
			} else {
				row = n - (next / n) - 1
			}

			if next%n == 0 {
				if (next/n)%2 == 0 {
					col = 0
				} else {
					col = n - 1
				}
			} else {
				if (next/n)%2 == 0 {
					col = next % n
				} else {
					col = n - (next % n)
				}
			}

			if board[row][col] == (n*n) || next == (n*n) {
				return distance
			} else if board[row][col] != -1 {
				queue = append(queue, board[row][col])
			} else {
				queue = append(queue, next)
			}
			fmt.Println(queue)
		}
		distance++
	}

	return -1
}
