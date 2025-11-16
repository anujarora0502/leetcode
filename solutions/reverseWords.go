package solutions

func ReverseWords(s string) string {
	s = trimWhiteSpaces(s)
	s = reverseString(s, 0, len(s)-1)
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " && end == -1 {
			start = i
			end = 0
		}else if string(s[i]) == " " && end == 0 {
			s = reverseString(s, start, i-1)
			end = -1
		}else if string(s[i]) == " "{
			s = s[:i] + s[i+1:]
			i--
		}
	}

	s = reverseString(s, start, len(s)-1)
	return s
}

func reverseString(s string, start int, end int) string {
	temp := ""
	i := start
	for i <= end {
		temp = string(s[i]) + temp
		i++
	}
	if start == 0 && end == len(s)-1 {
		return temp
	} else if end == len(s)-1 {
		s = s[:start] + temp
	} else if start == 0 {
		s = temp + s[end+1:]
	} else {
		s = s[:start] + temp + s[end+1:]
	}
	return s
}

func trimWhiteSpaces(s string) string {
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			s = s[i:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			s = s[:i+1]
			break
		}
	}

	return s
}
