package solutions

import "fmt"

func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		dp[i] = make([]int, len(s2))
	}
	return ilHelper("", s1, s2, s3, 0, 0, dp)
}

func ilHelper(ssf string, s1 string, s2 string, s3 string, i int, j int, dp [][]int) bool {
	fmt.Println(ssf)
	if ssf != s3[:len(ssf)] {
		return false
	}

	if i < len(s1) && j < len(s2) && dp[i][j] != 0 {
		if dp[i][j] == 1 {
			return true
		} else {
			return false
		}
	}

	if s3 == ssf {
		return true
	}

	result1 := ilHelper(ssf+string(s1[i]), s1, s2, s3, i+1, j, dp)

	if i+1 < len(s1) && j < len(s2) && result1 {
		if i+1 < len(s1) && j < len(s2) {
			dp[i+1][j] = 1
		}
		return true
	} else if i+1 < len(s1) && j < len(s2) && !result1 {
		dp[i+1][j] = -1
		return false
	}

	result2 := ilHelper(ssf+string(s2[j]), s1, s2, s3, i, j+1, dp)
	if i <= len(s1) && j+1 < len(s2) && result2 {
		dp[i][j] = 1
		return true
	} else if i <= len(s1) && j+1 <= len(s2) && !result2 {
		dp[i][j] = -1
		return false
	}
	return false
}
