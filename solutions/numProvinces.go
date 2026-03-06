package solutions

func findCircleNum(isConnected [][]int) int {
	isVisited := make([]bool, len(isConnected))

	count := 0

	for i := range isConnected {
		if !isVisited[i] {
			count++
			dfs(isConnected, i, isVisited)
		}
	}

	return count
}

func dfs(isConnected [][]int, v int, isVisited []bool) {
	isVisited[v] = true
	for i := range isConnected {
		if isConnected[v][i] == 1 && !isVisited[i] {
			dfs(isConnected, i, isVisited)
		}
	}
}
