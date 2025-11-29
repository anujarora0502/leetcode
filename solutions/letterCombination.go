package solutions

import (
	"fmt"
	"strconv"
)

func LetterCombinations(digits string) []string {

	letters := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	return combinationHelper("", digits, letters)

}

func combinationHelper(ssf string, digits string, letters [][]string) []string {
	if digits == "" {
		result := make([]string, 0)
		result = append(result, ssf)
		return result
	}

	firstDigit, _ := strconv.Atoi(string(digits[0]))
	index := firstDigit - 2
	fmt.Println(firstDigit)

	result := make([]string, 0)

	for _, e := range letters[index] {
		result = append(result, combinationHelper(ssf+e, digits[1:], letters)...)
	}

	return result
}
