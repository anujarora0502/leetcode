package solutions

func SimplifyPath(path string) string {
	var stack []byte

	for _, c := range path {
		if string(c) == "/" && len(stack) != 0 && string(stack[len(stack)-1]) == "/" {
			continue
		}
		if string(c) == "/" && len(stack) != 0 && string(stack[len(stack)-1]) == "." {
			lengthOfPeriods := 0
			for string(stack[len(stack)-lengthOfPeriods-1]) == "." {
				if lengthOfPeriods > 2 {
					break
				}
				lengthOfPeriods++
			}

			if string(stack[len(stack)-lengthOfPeriods-1]) != "/" {
				if string(stack[len(stack)-lengthOfPeriods-1]) != "." {
					stack = push(stack, byte(c))
					continue
				}
			}

			if lengthOfPeriods > 2 {
				stack = push(stack, byte(c))
			}

			if lengthOfPeriods == 1 {
				stack = stack[:len(stack)-1]
			}

			if lengthOfPeriods == 2 {
				for string(stack[len(stack)-1]) != "/" {
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 1 {
					continue
				}
				stack = stack[:len(stack)-1]
				for string(stack[len(stack)-1]) != "/" {
					stack = stack[:len(stack)-1]
				}
			}
			continue
		}
		stack = push(stack, byte(c))
	}

	if string(stack[len(stack)-1]) == "." {
		lengthOfPeriods := 0
		for string(stack[len(stack)-lengthOfPeriods-1]) == "." {
			if lengthOfPeriods > 2 {
				break
			}
			lengthOfPeriods++
		}

		if string(stack[len(stack)-lengthOfPeriods-1]) != "/" {
			if string(stack[len(stack)-lengthOfPeriods-1]) != "." {
				return string(stack)
			}
		}

		if lengthOfPeriods == 1 {
			stack = stack[:len(stack)-1]
		}

		if lengthOfPeriods == 2 {
			for string(stack[len(stack)-1]) != "/" {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 1 {
				return string(stack)
			}
			stack = stack[:len(stack)-1]
			for string(stack[len(stack)-1]) != "/" {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if string(stack[len(stack)-1]) == "/" && len(stack) > 1 {
		for string(stack[len(stack)-1]) == "/" {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

func push(stack []byte, element byte) []byte {
	stack = append(stack, element)
	return stack
}
