package solutions

func Permute(nums []int) [][]int {
	visited := make([]bool, len(nums))
	arrSoFar := make([]int, 0)

	result := make([][]int, 0)
	permuteHelper(nums, visited, arrSoFar, &result)

	return result
}

func permuteHelper(nums []int, visited []bool, arrSoFar []int, result *[][]int) {
	if len(arrSoFar) == len(nums) {
		tempArr := make([]int, len(arrSoFar))
		copy(tempArr, arrSoFar)
		*result = append(*result, tempArr)
		return
	}

	for i, e := range nums {
		if !visited[i] {
			visited[i] = true
			arrSoFar = append(arrSoFar, e)
			permuteHelper(nums, visited, arrSoFar, result)
			arrSoFar = arrSoFar[:len(arrSoFar)-1]
			visited[i] = false
		}
	}
}
