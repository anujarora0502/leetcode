package solutions

import (
	"fmt"
	"strconv"
)

func Calculate(s string) int {
	stack := make([]string, 0)

	for i := 0; i < len(s); i++ {
		fmt.Println(stack)
		e := string(s[i])
		element := string(e)
		if element == "(" {
			stack = append(stack, element)
		} else if element == ")" {
			if len(stack) <= 3 {
				continue
			}
			number2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			operator := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			number1, _ := strconv.Atoi(stack[len(stack)-1])
			if operator == "+" {
				stack[len(stack)-1] = strconv.Itoa(number1 + number2)
			} else {
				stack[len(stack)-1] = strconv.Itoa(number1 - number2)
			}
		} else if element == " " {
			continue
		} else if element == "-" || element == "+" {
			stack = append(stack, element)
		} else {
			if len(stack) == 0 {
				stack = append(stack, element)
			} else {
				top := stack[len(stack)-1]
				if top == "(" {
					stack = append(stack, element)
				} else if top == "+" || top == "-" {
					if len(stack) <= 1 {
						stack[len(stack)-1] += element
						continue
					}
					beforeTop := stack[len(stack)-2]
					if beforeTop == "(" {
						stack[len(stack)-1] += element
						continue
					}
					number1, _ := strconv.Atoi(beforeTop)
					number2, _ := strconv.Atoi(element)
					stack = stack[:len(stack)-1]
					if top == "+" {
						stack[len(stack)-1] = strconv.Itoa(number1 + number2)
					} else {
						stack[len(stack)-1] = strconv.Itoa(number1 - number2)
					}
				} else {
					stack[len(stack)-1] += element
				}
			}
		}
	}

	fmt.Println(stack)

	var result int

	if len(stack) == 3 {
		if stack[0] == "-" {
			result, _ = strconv.Atoi(stack[0] + stack[2])
			return result
		}
	}

	result, _ = strconv.Atoi(stack[len(stack)-1])

	return result
}
