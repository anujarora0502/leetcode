package solutions

import (
	"fmt"
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	p1 := 0
	p2 := len(s) - 1
	s = strings.ToLower(s)

	for p1 < p2 {
		if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(string(s[p1])) {
			p1++
			continue
		}
		if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(string(s[p2])) {
			p2--
			continue
		}

		if s[p1] != s[p2] {
			return false
		}
		fmt.Println(string(s[p1]))
		fmt.Println(string(s[p2]))
		p1++
		p2--
	}

	return true
}
