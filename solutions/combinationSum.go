package solutions

func CombinationSum(candidates []int, target int) [][]int {
	return helper(candidates, target, make([]int, 0), 0)
}

func helper(candidates []int, target int, arrSoFar []int, index int) [][]int {
	if target < 0 {
		return nil
	}

	if target == 0 {
		result := make([][]int, 0)
		arr := make([]int, len(arrSoFar))
		copy(arr, arrSoFar)
		result = append(result, arr)
		return result
	}

	result := make([][]int, 0)

	for i := index; i < len(candidates); i++ {
		temp := helper(candidates, target-candidates[i], append(arrSoFar, candidates[i]), i)
		if temp != nil {
			result = append(result, temp...)
		}
	}

	return result
}
