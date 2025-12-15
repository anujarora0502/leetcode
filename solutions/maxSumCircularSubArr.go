package solutions

func MaxSubarraySumCircular(nums []int) int {
	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {
		j := i
		var end int
		if j == 0 {
			end = len(nums) - 1
		} else {
			end = j - 1
		}
		currSum := 0
		for {
			currSum += nums[j]
			if maxSum < currSum {
				maxSum = currSum
			}
			if currSum < 0 {
				currSum = 0
			}
			if j == end {
				break
			}
			j++
			if j == len(nums) {
				j = 0
			}
		}

	}

	return maxSum
}
